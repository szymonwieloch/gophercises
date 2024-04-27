package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := parseConfig()

	//mux := http.NewServeMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	urlsWithMap := mapHandler(pathsToUrls, http.HandlerFunc(defaultHandler))

	urlsWithYAML, err := yamlHandler(cfg.yamlPath, urlsWithMap)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting server on port ", cfg.port)
	addr := fmt.Sprintf(":%d", cfg.port)

	http.ListenAndServe(addr, urlsWithYAML)

}
