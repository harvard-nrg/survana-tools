package main

import (
	"text/template"
)

var tpl_group = `
[[AdvancedChoices]]
{{ range $index, $item := .Answer.SItems }}
[[Choice:{{ $item.Value}}]]
{{ $item.Html }}{{ end }}
`

func init() {
	template.Must(Templates.New("group").Parse(tpl_group))
}
