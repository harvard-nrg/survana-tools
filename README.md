survana-tools
=============

A set of tools for Survana.

export-qualtrics
----------------
Exports Survana forms to Qualtrics TXT format suitable for import.

* Prerequisites: Go 1.x
* Testing: Tested on OS X.

Installation

```
cd export-qualtrics
make
./bin/export-qualtrics -i file.json -o file.txt
```

Usage:

* `-i` - input JSON file containing 1 form, as exported from the database
* `-o` - output file in Qualtrics TXT format
* `-f` - overwrite output file if it exists
* `-g` - output debug information
* '--help' - output the help page

To import the TXT file into Qualtrics, follow this guide: http://www.qualtrics.com/university/researchsuite/advanced-building/advanced-options-drop-down/import-and-export-surveys/


export-responses
----------------

Prerequisites: easy_install xlsxwriter
