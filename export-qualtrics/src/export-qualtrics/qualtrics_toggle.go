package main

import (
	"text/template"
)

var tpl_toggle = `
[[AdvancedChoices]]{{ range $index, $item := .Answer.SItems }}
[[Choice]]
{{ $item.Html }}{{ end }}
`

func init() {
	template.Must(Templates.New("toggle").Parse(tpl_toggle))
}
