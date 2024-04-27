package main

import "flag"

type config struct {
	port     uint16
	yamlPath string
}

func parseConfig() config {
	port := flag.Uint64("port", 8080, "Port to launch HTTP server on")
	yamlPath := flag.String("path", "urls.yaml", "Path to the YAML file with shortened URLs")
	return config{
		port:     uint16(*port),
		yamlPath: *yamlPath,
	}
}
