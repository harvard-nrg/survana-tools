package main

import (
	"text/template"
)

var tpl_box = `
`

func init() {
	template.Must(Templates.New("box").Parse(tpl_box))
}
