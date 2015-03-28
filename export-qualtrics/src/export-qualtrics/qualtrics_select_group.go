package main

import (
	"text/template"
)

var tpl_select_group = `
[[AdvancedChoices]]
{{ range $index, $item := .Answer.SItems }}
[[Choice:{{ $item.Value}}]]
{{ $item.Html }}{{ end }}
`

func init() {
	template.Must(Templates.New("select_group").Parse(tpl_select_group))
}
