package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type option struct {
	Arc  string
	Text string
}

type chapter struct {
	Title   string
	Story   []string
	Options []option
}

type story map[string]chapter

func parseStory(path string) (story, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %w", path, err)
	}
	var s story
	err = json.Unmarshal(f, &s)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON file %s: %w", path, err)
	}
	return s, nil
}
