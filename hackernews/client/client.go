package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const hnAPI = "https://hacker-news.firebaseio.com/v0"

type StoryID int

func TopStories() ([]StoryID, error) {
	resp, err := http.Get(hnAPI + "/topstories.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	var result []StoryID
	err = dec.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func GetStory(id StoryID) (Story, error) {
	url := fmt.Sprintf("%s/item/%d.json", hnAPI, id)
	resp, err := http.Get(url)
	if err != nil {
		return Story{}, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	var story Story
	err = dec.Decode(&story)
	if err != nil {
		return Story{}, err
	}
	return story, err
}

type Story struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`

	// Only one of these should exist
	Text string `json:"text"`
	URL  string `json:"url"`
}
