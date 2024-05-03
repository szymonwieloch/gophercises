package main

import "flag"

// Commanline configuratio of the application
type config struct {
	url    string
	output string
	depth  uint
}

// Parses command line configuratio
func parseConfig() config {
	url := flag.String("url", "", "URL to start reading pages from")
	output := flag.String("output", "sitemap.xml", "Path to the output XML file where the site map is stored")
	depth := flag.Uint("depth", 0, "Depth for looking into subpages")
	flag.Parse()
	cfg := config{
		url:    *url,
		output: *output,
		depth:  *depth,
	}

	return cfg
}
