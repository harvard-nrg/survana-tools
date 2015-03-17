package main

import (
	"text/template"
)

var tpl_radio_group = `
[[AdvancedChoices]]
{{ range $index, $item := .Answer.SItems.Items }}
[[Choice:{{ $item.Value}}]]
{{ $item.Key }}{{ end }}
`

func init() {
	template.Must(Templates.New("radio_group").Parse(tpl_radio_group))
}
