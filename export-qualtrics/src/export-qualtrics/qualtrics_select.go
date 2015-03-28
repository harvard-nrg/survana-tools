package main

import (
	"text/template"
)

var tpl_select = `
[[AdvancedChoices]]
{{ range $index, $item := .Answer.SItems }}{{ if eq $item.SType "optgroup" }}{{ range $optindex, $optitem := $item.SItems }}
[[Choice]]
{{ $optitem.Html }}{{end}}{{ else }}
[[Choice]]
{{ $item.Html }}{{ end }}{{ end }}
`

func init() {
	template.Must(Templates.New("select").Parse(tpl_select))
}
