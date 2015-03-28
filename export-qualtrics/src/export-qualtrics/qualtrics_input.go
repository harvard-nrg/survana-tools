package main

import (
	"text/template"
)

var tpl_input = `
`

func init() {
	template.Must(Templates.New("input").Parse(tpl_input))
}
