package main

import (
	"flag"
	"time"
)

// Configuration of the application
type config struct {
	csvPath string
	shuffle bool
	limit   time.Duration
}

// Parses command line configuration
func parseConfig() config {
	csvPath := flag.String("p", "problems.csv", "Path to the CSV with problems")
	shuffle := flag.Bool("s", false, "Shuffle questions")
	limit := flag.Uint("l", 0, "Time limit in seconds for the quiz, 0 = no limit ")
	flag.Parse()
	return config{
		csvPath: *csvPath,
		shuffle: *shuffle,
		limit:   time.Second * time.Duration(*limit),
	}
}
