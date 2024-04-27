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

type arc struct {
	Title   string
	Story   []string
	Options []option
}

func parseStory(path string) (map[string]arc, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %w", path, err)
	}
	var story map[string]arc
	err = json.Unmarshal(f, &story)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON file %s: %w", path, err)
	}
	return story, nil
}
