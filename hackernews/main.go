package main

import (
	_ "embed"
	"log"
	"net/http"
	"net/url"
	"slices"
	"strings"
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

const wantStories = 30

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	stories, err := getTopStories(wantStories)
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

func takeN[T any](in []T, n int) ([]T, []T) {
	take := min(n, len(in))
	return in[:take], in[take:]
}

type asyncResult struct {
	story client.Story
	err   error
}

func getBatch(storyIDs []client.StoryID) []client.Story {
	ch := make(chan asyncResult)
	for _, storyId := range storyIDs {
		go func(storyId client.StoryID) {
			story, err := client.GetStory(storyId)
			ch <- asyncResult{
				story: story,
				err:   err,
			}
		}(storyId)
	}
	var stories []client.Story
	for range storyIDs {
		ar := <-ch
		if ar.err != nil {
			log.Println("Error getting story: ", ar.err)
			continue
		}
		stories = append(stories, ar.story)
	}
	return stories
}

func getTopStories(count int) ([]client.Story, error) {
	topStories, err := client.TopStories()
	if err != nil {
		return nil, err
	}
	remaining := topStories
	var stories []client.Story
	for len(stories) < count && len(remaining) > 0 {
		missing := count - len(stories)
		batchSize := (missing * 5 / 4) + 3
		var batchIDs []client.StoryID
		batchIDs, remaining = takeN(remaining, batchSize)
		batchStories := getBatch(batchIDs)
		for _, story := range batchStories {
			if isStoryLink(story) {
				stories = append(stories, story)
			}
		}
	}
	// we may have too many stories
	stories = stories[:min(count, len(stories))]
	//sort
	storyIDToIdx := map[client.StoryID]int{}
	for idx, storyId := range topStories {
		storyIDToIdx[storyId] = idx
	}
	slices.SortFunc(stories, func(a, b client.Story) int {
		return storyIDToIdx[a.ID] - storyIDToIdx[b.ID]
	})
	return stories, nil
}

func isStoryLink(story client.Story) bool {
	return story.Type == "story" && story.URL != ""
}

func main() {
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
