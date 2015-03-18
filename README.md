survana-tools
=============

A set of tools for Survana.

export-qualtrics
----------------
In progress.

```
cd export-qualtrics
make
./bin/export-qualtrics -i file.json -o file.txt
```

Usage:

* `-i` - input JSON file containing 1 form, as exported from the database
* '-o' - output file in Qualtrics TXT format
* '-f' - overwrite output file if it exists

export-responses
----------------

Prerequisites: easy_install xlsxwriter
