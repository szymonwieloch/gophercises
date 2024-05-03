package main

import (
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

//go:embed templates/layout.html
var templateData embed.FS

func serve(s story, port uint16) {
	hdl := handler{
		tmpl: template.Must(template.ParseFS(templateData, "templates/layout.html")),
		str:  s,
	}
	addr := fmt.Sprintf(":%d", port)
	fmt.Println("Listening on address ", addr)
	http.ListenAndServe(addr, &hdl)
}

type handler struct {
	tmpl *template.Template
	str  story
}

func (hdl *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	storyName := strings.TrimLeft(req.URL.Path, "/")
	if storyName == "" {
		storyName = "intro"
	}
	chapter, ok := hdl.str[storyName]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Chapter not found"))
		return
	}
	err := hdl.tmpl.Execute(w, chapter)
	if err != nil {
		log.Fatal(err)
	}
}
