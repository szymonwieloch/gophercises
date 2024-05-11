package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/szymonwieloch/gophercises/linkparser"
)

// Cleans the link or returns empty string if link does not match the criteria
func cleanURL(link, base *url.URL) string {
	if !link.IsAbs() {
		link = base.ResolveReference(link)
	}
	if link.Host != base.Host {
		return ""
	}
	if link.Scheme != "http" && link.Scheme != "https" {
		return ""
	}
	link.Fragment = ""
	link.Path = strings.TrimRight(link.Path, "/")
	return link.String()
}

// Parses the given URL path and outputs all of found links that belong to the same host
func pollSingleURL(urlPath string) []string {
	resp, err := http.Get(urlPath)
	if err != nil {
		log.Printf("Error getting URL '%s': %v", urlPath, err)
		return nil
	}
	defer resp.Body.Close()

	originalUrl, err := url.Parse(urlPath)
	if err != nil {
		// links should be verified before getting here
		log.Fatal(err)
	}

	// possibly redirected to a different URL or site
	if originalUrl.Host != resp.Request.URL.Host {
		// we got redirected, skip all the links
		return nil
	}
	links, err := linkparser.ParseReader(resp.Body)
	if err != nil {
		log.Printf("Error parsing HTML for URL '%s': %v", urlPath, err)
		return nil
	}

	result := []string{}
	for _, link := range links {
		linkURL, err := url.Parse(link.Href)
		if err != nil {
			log.Printf("Error parsing link '%s': %v", link.Href, err)
			continue
		}
		cleaned := cleanURL(linkURL, resp.Request.URL)
		if cleaned == "" {
			continue
		}
		result = append(result, cleaned)
	}
	return result
}

// Parses the whole site by following internal links up to the given depth
func pollSite(urlPath string, depth uint) []string {
	knownLinks := map[string]bool{
		urlPath: true,
	}
	toDo := []string{urlPath}

	for range depth {
		newToDo := []string{}
		for _, link := range toDo {
			newLinks := pollSingleURL(link)
			for _, nl := range newLinks {
				if knownLinks[nl] {
					continue
				}
				knownLinks[nl] = true
				newToDo = append(newToDo, nl)
			}
		}
		toDo = newToDo
	}
	result := make([]string, 0, len(knownLinks))
	for link := range knownLinks {
		result = append(result, link)
	}
	return result
}
