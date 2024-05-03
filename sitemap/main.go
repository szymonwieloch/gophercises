package main

import (
	"fmt"
	"log"
	"net/url"
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

	pullSite(cfg.url)

}
