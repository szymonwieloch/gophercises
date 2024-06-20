package main

import (
	_ "embed"
	"html/template"
)

//go:embed index.htm
var indexPage []byte

//go:embed view.htm
var viewPage string

//go:embed choice.htm
var choicePage string

var viewTemplate *template.Template
var choiceTemplate *template.Template

type viewChoiceData struct {
	Name       string
	ChoiceLink string
}

type viewData struct {
	ImgLink string
	Choices []viewChoiceData
}

type choiceData struct {
	Choice  string
	Options []choiceOptionData
}

type choiceOptionData struct {
	ImgLink  string
	ViewLink string
}

func init() {
	viewTemplate = template.Must(template.New("view").Parse(viewPage))
	choiceTemplate = template.Must(template.New("choice").Parse(choicePage))
}
