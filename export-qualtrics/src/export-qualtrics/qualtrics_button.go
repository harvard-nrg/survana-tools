package main

import (
	"text/template"
)

var tpl_button = `
`

func init() {
	template.Must(Templates.New("button").Parse(tpl_button))
}
