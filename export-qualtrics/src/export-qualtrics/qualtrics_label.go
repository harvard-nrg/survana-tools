package main

import (
	"text/template"
)

var tpl_label = `
[[Question:DB]]
{{ .Answer.Html }}
`

func init() {
	template.Must(Templates.New("label").Parse(tpl_label))
}
