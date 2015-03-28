package main

import (
	"text/template"
)

var tpl_slider = `
`

func init() {
	template.Must(Templates.New("slider").Parse(tpl_slider))
}
