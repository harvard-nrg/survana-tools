package main

import (
	"text/template"
)

var tpl_link = `
`

func init() {
	template.Must(Templates.New("link").Parse(tpl_link))
}
