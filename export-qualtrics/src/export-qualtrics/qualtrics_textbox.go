package main

import (
	"text/template"
)

var tpl_textbox = `
`

func init() {
	template.Must(Templates.New("textbox").Parse(tpl_textbox))
}
