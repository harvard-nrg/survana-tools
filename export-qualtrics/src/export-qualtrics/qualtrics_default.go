package main

import (
        "text/template"
       )

var tpl_default = `
default
`

func init() {
    template.Must(Templates.New("default").Parse(tpl_default))
}
