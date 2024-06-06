package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/szymonwieloch/gophercises/hackernews/client"
)

//go:embed index.htm
var templateStr string
var templates *template.Template

func init() {
	templates = template.Must(template.New("main").Parse(templateStr))
}

func handler(w http.ResponseWriter, r *http.Request, cache cacheStrategy) {
	start := time.Now()
	stories, err := cache.getTopStories()
	if err != nil {
		http.Error(w, "Failed to load top stories", http.StatusInternalServerError)
		return
	}

	var items []item
	for _, story := range stories {
		var host string
		u, err := url.Parse(story.URL)
		if err == nil {
			host = strings.TrimPrefix(u.Hostname(), "www.")
		}
		items = append(items, item{Story: story, Host: host})
	}
	td := templateData{
		Stories: items,
		Time:    time.Since(start),
	}
	err = templates.Execute(w, &td)
	if err != nil {
		log.Fatal(err)
	}

}

func createHandler(a args) http.HandlerFunc {
	var cache cacheStrategy
	switch a.Cache {
	case useCacheNone:
		cache = cacheNone{
			count: int(a.Count),
		}
	case useCacheRefres:
		cache = &cacheRefresh{
			count:  int(a.Count),
			period: a.Period,
		}
	case useCacheBackground:
		cache = newBackgroundCache(int(a.Count), a.Period)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, cache)
	}
}

func main() {
	var a args
	arg.MustParse(&a)
	http.HandleFunc("/", createHandler(a))

	// Start the server
	port := fmt.Sprintf(":%d", a.Port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// item is the same as the hn.Item, but adds the Host field
type item struct {
	client.Story
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}
