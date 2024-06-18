package main

import (
	_ "embed"
	"html/template"
)

//go:embed index.htm
var indexPage []byte

//go:embed view.htm
var viewPage string

var viewTemplate *template.Template

type viewData struct {
	ImagePath string
	Options   []string
}

func init() {
	viewTemplate = template.Must(template.New("view").Parse(viewPage))
}
