package main

import (
	"text/template"
)

var tpl_question = `
[[Question:{{ .QualtricsType }}]]
[[ID:{{ .Answer.SId }}]]
{{ .Question.Html }}`

func init() {
	template.Must(Templates.New("question").Parse(tpl_question))
}
