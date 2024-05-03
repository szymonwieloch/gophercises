package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/szymonwieloch/gophercises/linkparser"
)

func pullSite(urlPath string) {
	linkparser.ParseReader(strings.NewReader(""))
	resp, err := http.Get(urlPath)
	if err != nil {
		log.Printf("Error getting URL '%s': %v", urlPath, err)
		return
	}
	defer resp.Body.Close()
	links, err := linkparser.ParseReader(resp.Body)
	if err != nil {
		log.Printf("Error parsing HTML for URL '%s': %v", urlPath, err)
		return
	}
	base, err := url.Parse(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, link := range links {
		linkUrl, err := url.Parse(link.Href)
		if err != nil {
			log.Printf("Error parsing link '%s': %v", link.Href, err)
			continue
		}
		if !linkUrl.IsAbs() {
			linkUrl = linkUrl.ResolveReference(base)
		}
		fmt.Println(linkUrl)
	}
}
