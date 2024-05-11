package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Site map builder")
	cfg := parseConfig()
	if cfg.url == "" {
		log.Fatal("Please provide the initial URL")
	}
	_, err := url.Parse(cfg.url)
	if err != nil {
		log.Fatal(err)
	}
	// clean the input url
	cfg.url = strings.TrimRight(cfg.url, "/")

	links := pollSite(cfg.url, cfg.depth)
	saveSitemap(links, cfg.output)
}
