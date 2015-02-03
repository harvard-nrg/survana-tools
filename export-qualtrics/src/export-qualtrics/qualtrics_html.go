package main

import (
        "text/template"
       )

var tpl_html = `
html
`

func init() {
    template.Must(Templates.New("html").Parse(tpl_html))
}
