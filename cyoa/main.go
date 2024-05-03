package main

import (
	"log"
)

func main() {
	cfg := parseConfig()

	str, err := parseStory(cfg.storyFile)
	if err != nil {
		log.Fatal(err)
	}
	serve(str, cfg.port)

}
