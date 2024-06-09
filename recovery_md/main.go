package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func regular(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Regular"))
}

func panics(w http.ResponseWriter, r *http.Request) {
	panic("Something went wrong!")
}

func latePanic(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Some initial text"))
	panic("Something went terribly wrong!")
}

func showSource(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	line := r.URL.Query().Get(("line"))

	lineInt := -1
	var err error
	if line != "" {
		lineInt, err = strconv.Atoi(line)
		if err != nil {
			lineInt = -1
			log.Println("Could not parse line ", line)
		}
	}
	fmt.Println("show file ", path, ", line", lineInt)
	file, err := os.Open(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(([]byte("Could not open file")))
		log.Println("Could not open file ", path)
		return
	}
	defer file.Close()
	sourceCode, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(([]byte("Could not open file")))
		log.Println("Could not read file ", path)
		return
	}
	//w.Write(sourceCode)
	lexer := lexers.Get("go")
	if lexer == nil {
		lexer = lexers.Fallback
	}
	style := styles.Get("github")
	if style == nil {
		style = styles.Fallback
	}
	htmlOptions := []html.Option{html.TabWidth(4), html.WithLineNumbers(true)}
	if lineInt != -1 {
		hLines := [][2]int{[2]int{lineInt, lineInt}}
		htmlOptions = append(htmlOptions, html.HighlightLines(hLines))
	}

	//formatter := formatters.Get("html")
	formatter := html.New(htmlOptions...)
	iterator, err := lexer.Tokenise(nil, string(sourceCode))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not tokenise file"))
		log.Println("Could not tokenise ", path)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = formatter.Format(w, style, iterator)
	if err != nil {
		log.Println("Error when formatting code: ", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/regular", regular)
	mux.HandleFunc("/panics", panics)
	mux.HandleFunc("/latepanic", latePanic)
	mux.HandleFunc("/source", showSource)
	log.Fatal(http.ListenAndServe(":3000", recoveryMw(mux, true)))
}
