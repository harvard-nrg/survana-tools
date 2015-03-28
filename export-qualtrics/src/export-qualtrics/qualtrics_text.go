package main

import (
	"text/template"
)

var tpl_text = `
`

func init() {
	template.Must(Templates.New("text").Parse(tpl_text))
}
