# find_script

This tool is intended to find fields in a Postgres database that contain a `<script>` tag.

Pen testers like to try and break stuff, one method is to see if fields can have javascript embedded.

## Usage

```bash
DATABASE_URL="heroku format" find_script
```

Output to be determined.

## Local environment

To save having to include the DATABASE_URL in the commandline, copy the `example.env` file to `.env` and set the database that you want to examine.
