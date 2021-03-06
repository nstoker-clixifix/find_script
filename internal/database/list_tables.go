package database

import (
	"database/sql"
	"fmt"
	"strings"

	"log"
)

func ScanTables() {
	rows, err := DB.Query(`SELECT table_name FROM information_schema.tables WHERE table_schema='public' and table_name not like 'pg_%' and table_name != 'schema_migrations'`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			log.Fatal(err)
		}

		checkFields(table)
	}
}

func checkFields(table string) {
	fields := getTextFields(table)
	if len(fields) == 0 {
		return
	}

	field_list := strings.Join(fields[:], ",")
	where_clause := buildWhereClause(fields)
	query := "SELECT "
	if table == "linked_signatures" {
		query += "signature_id"
	} else {
		query += "id"
	}
	query += fmt.Sprintf(", %s FROM %s WHERE %s", field_list, table, where_clause)
	rows, err := DB.Queryx(query)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("checkFields: %v (%s)", err, query)
	}
	defer rows.Close()

	for rows.Next() {
		values, err := rows.SliceScan()
		if err != nil {
			log.Printf("%s - %v", table, err)
			// log.Fatal(err)
		}
		if values != nil {
			// log.Printf("%s row %+v", table, values[0])
			listOffences(table, fields, values)
		}
	}
}

func listOffences(table string, fields []string, values []interface{}) {
	var id int64
	var ok bool
	for column, value := range values {
		if column == 0 {
			id, ok = value.(int64)
			if !ok {
				log.Printf("%s.%s error parsing id for '%s'", table, fields[:0], value)
				return
			}
			continue
		}
		if value == nil {
			continue
		}

		matchedText := findScriptText(value)
		if matchedText != "" {
			log.Printf("%s row %d '%s'", table, id, matchedText)
		}
	}
}

const MAX_WIDTH = 60

func findScriptText(value interface{}) string {
	text := fmt.Sprintf("%v", value)
	startPosition := strings.Index(text, "<script")
	if startPosition == -1 {
		return ""
	}
	closePosition := strings.Index(text, "</script>") + 9
	if closePosition == -1 {
		closePosition = startPosition + MAX_WIDTH
	}

	if closePosition > startPosition+MAX_WIDTH {
		closePosition = startPosition + MAX_WIDTH
	}
	if closePosition > len(text) {
		return text[startPosition:]
	}

	return text[startPosition:closePosition]
}

func buildWhereClause(fields []string) string {
	whereClause := []string{}

	for _, field := range fields {
		whereClause = append(whereClause, constraint(field))
	}
	return strings.Join(whereClause[:], " OR ")
}

func constraint(field string) string {
	cons := fmt.Sprintf("LOWER(%s) like '<", field)
	cons += "%%script%%>'"
	return cons
}
func getTextFields(table string) []string {
	fields := []string{}
	rows, err := DB.Query(`SELECT column_name, data_type FROM information_schema.columns WHERE table_name=$1`, table)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var column string
		var dataType string
		if err = rows.Scan(&column, &dataType); err != nil {
			log.Fatal(err)
		}

		switch dataType {
		case "bigint", "double precision", "int", "integer", "numeric":
		case "boolean":
		case "date", "timestamp", "timestamp with time zone", "timestamp without time zone":
		case "json", "jsonb":
		case "uuid":
		case "text", "character varying", "name":
			fields = append(fields, column)
		case "ARRAY", "USER-DEFINED":
		default:
			log.Printf("unhandled data type '%s'", dataType)
		}
	}
	return fields
}
