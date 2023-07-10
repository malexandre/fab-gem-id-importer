# CSV importer into a GEM event

Tool to import users from a CSV to a GEM event in Flesh and Blood.

CSV should have named columns, with at least the column `gemid`. If `name` is also a column, it will be used to display errors when users are not found.
If not, the GEM Id alone will be displayed.

The file extension is used to determine the sperators: Commas (.csv), Tabs (.tsv) or semicolon (.ssv).

## Usage

You can chose to set the `username` and `password` either through env vars (via .env or through the command line) or through flags like the following example:

```bash
fab-gem-id-importer -u username -p password -e 42314 ./path-to-file.csv
```

At any time, you can use the `-h` flag:

```bash
$ fab-gem-id-importer -h
fab-gem-id-importer is a tool to import users from a CSV to a GEM event. User login and password can be set through .env or with flags.

Usage:
  fab-gem-id-importer [flags] path_to_file.(csv|tsv|ssv)

Flags:
  -e, --event-id string   GEM event ID to update with the users. Found in the GEM URL of the event.
  -h, --help              help for fab-gem-id-importer
  -p, --password string   Password of the GEM account to log in as
  -u, --username string   Name of the GEM account to log in as
```

## Next step

I should manage release executables in the following days.
