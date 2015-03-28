package main

import (
	"text/template"
)

var tpl_checkbox_group = `
[[AdvancedChoices]]
{{ range $index, $item := .Answer.SItems }}
[[Choice:{{ $item.Value}}]]
{{ $item.Html }}{{ end }}
`

func init() {
	template.Must(Templates.New("checkbox_group").Parse(tpl_checkbox_group))
}
