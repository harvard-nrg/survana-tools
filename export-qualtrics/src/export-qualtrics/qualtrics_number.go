package main

import (
	"text/template"
)

var tpl_number = `
`

func init() {
	template.Must(Templates.New("number").Parse(tpl_number))
}
