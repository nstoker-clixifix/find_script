package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

func ScanTables() {
	rows, err := DB.Query(`SELECT table_name FROM information_schema.tables WHERE table_schema='PUBLIC'`)
	if err != nil {
		logrus.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var table string

		err = rows.Scan(&table)
		if err != nil {
			logrus.Fatal(err)
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
	query := fmt.Sprintf("SELECT id, %s FROM %s WHERE %s", field_list, table, where_clause)
	rows, err := DB.Query(query)
	if err != nil && err != sql.ErrNoRows {
		logrus.Fatalf("checkFields: %v (%s)", err, query)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		values := make([]string, len(field_list))

		if err := rows.Scan(&id, &values); err != nil {
			logrus.Fatal(err)
		}

		logrus.Warnf("%s %d: %s", table, id, values)

	}
	logrus.Infof("%s %+v", table, fields)
}

func buildWhereClause(fields []string) string {
	whereClause := []string{}

	for _, field := range fields {
		whereClause = append(whereClause, constraint(field))
	}
	return strings.Join(whereClause[:], " AND ")
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
		logrus.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var column string
		var dataType string
		if err = rows.Scan(&column, &dataType); err != nil {
			logrus.Fatal(err)
		}

		switch dataType {
		case "bigint", "double precision", "int", "integer", "numeric":
		case "boolean":
		case "date", "timestamp", "timestamp with time zone", "timestamp without time zone":
		case "json", "jsonb":
		case "text", "character varying", "name", "uuid":
			fields = append(fields, column)
		case "ARRAY", "USER-DEFINED":
		default:
			logrus.Warnf("unhandled data type '%s'", dataType)
		}
	}
	return fields
}
