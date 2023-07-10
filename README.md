# fab-gem-id-importer

Tool to import users from a CSV to a GEM event in Flesh and Blood.

CSV should have named columns, with at least the column `gemid`. If `name` is also a column, it will be used to display errors when users are not found.
If not, the GEM Id alone will be displayed.

The file extension is used to determine the sperators: Commas (.csv), Tabs (.tsv) or semicolon (.ssv).
