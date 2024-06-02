package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/szymonwieloch/gophercises/hackernews/client"
)

//go:embed index.htm
var templateStr string
var templates *template.Template

func init() {
	templates = template.Must(template.New("main").Parse(templateStr))
}

func handler(w http.ResponseWriter, r *http.Request) {
	topStories, err := client.TopStories()
	if err != nil {
		http.Error(w, "Failed to load top stories", http.StatusInternalServerError)
		return
	}
	detailedStories := topStories[:30]
	var stories []client.Story
	for _, ds := range detailedStories {
		story, err := client.GetStory(ds)
		if err != nil {
			http.Error(w, "Failed to load top stories", http.StatusInternalServerError)
			return
		}
		stories = append(stories, story)
	}

	var items []item
	for _, story := range stories {
		items = append(items, item{Story: story, Host: "google.com"})
	}
	td := templateData{
		Stories: items,
		Time:    time.Second * 5,
	}
	err = templates.Execute(w, &td)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	stories, err := client.TopStories()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stories)

	http.HandleFunc("/", handler)

	// Start the server
	log.Fatal(http.ListenAndServe(":3000", nil))
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
