package main

import (
	"text/template"
)

var tpl_rtf = `
[[Question:DB]]
{{ .Answer.Html }}
`

func init() {
	template.Must(Templates.New("rtf").Parse(tpl_rtf))
}
