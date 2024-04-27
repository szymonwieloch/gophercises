package main

import "flag"

type config struct {
	port      uint16
	storyFile string
}

func parseConfig() config {
	port := flag.Uint64("port", 8080, "Port to open HTTP server on")
	storyFile := flag.String("story", "gopher.json", "Story file path")
	flag.Parse()

	return config{
		port:      uint16(*port),
		storyFile: *storyFile,
	}
}
