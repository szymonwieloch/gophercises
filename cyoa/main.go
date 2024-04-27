package main

import (
	"fmt"
	"log"
)

func main() {
	cfg := parseConfig()

	story, err := parseStory(cfg.storyFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(story)
}
