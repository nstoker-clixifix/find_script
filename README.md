# find_script

This tool is intended to find fields in a Postgres database that contain a `<script>` tag.

Pen testers like to try and break stuff, one method is to see if fields can have javascript embedded.

## Usage

```bash
Usage of ./findScriptTags:
  -d string
    	the heroku style database url for your database
  -e	looks in the environment for 'database_url' (this can be in a .env file 
  -help
    	displays the help text
```

Sample output:

```log
2021/08/06 08:56:15 Initialising
2021/08/06 08:56:15 connecting to database
2021/08/06 08:56:15 connected to de759a0055jpa7
2021/08/06 08:56:17 attachments row 3847 '<script>alert('Forename!')</script>'
...
2021/08/06 08:57:53 tickets row 66692 '<script>alert('Forename!')</script>'
2021/08/06 08:57:54 closing connection to de759a0055jpa7
```

## Local environment

To save having to include the DATABASE_URL in the commandline, copy the `example.env` file to `.env` and set the database that you want to examine.
