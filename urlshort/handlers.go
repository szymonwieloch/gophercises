package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Error: handler for %s not found", req.URL.Path)

}

func mapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		destination, ok := pathsToUrls[req.URL.Path]
		if !ok {
			fallback.ServeHTTP(w, req)
			return
		}
		http.Redirect(w, req, destination, http.StatusSeeOther)
	}
}

type mapping struct {
	URL  string
	Path string
}

func yamlHandler(pathToYaml string, fallback http.Handler) (http.HandlerFunc, error) {
	f, err := os.ReadFile(pathToYaml)
	if err != nil {
		return nil, fmt.Errorf("failed to open YAM file %v: %w", pathToYaml, err)
	}

	var mappings []mapping
	err = yaml.UnmarshalStrict(f, &mappings)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the YAML file %s: %w", pathToYaml, err)
	}

	urlToPath := map[string]string{}
	for _, mp := range mappings {
		urlToPath[mp.URL] = mp.Path
	}

	return mapHandler(urlToPath, fallback), nil
}
