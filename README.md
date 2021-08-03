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
INFO[0000] connecting to database                       
INFO[0000] connected to d285rc9qqf71p4                  
WARN[0001] attachments row 3847 '<script>alert('Forename!')</script>' 
INFO[0017] closing connection to %sd285rc9qqf71p4  
```

## Local environment

To save having to include the DATABASE_URL in the commandline, copy the `example.env` file to `.env` and set the database that you want to examine.
