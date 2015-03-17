package main

import (
	"text/template"
)

var tpl_instruction = `
[[Question:DB]]
{{ .Question.Html }}
`

func init() {
	template.Must(Templates.New("instruction").Parse(tpl_instruction))
}
