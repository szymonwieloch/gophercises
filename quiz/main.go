package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	cfg := parseConfig()
	problems, err := parseProblemsFile(cfg.csvPath)
	if err != nil {
		log.Fatal(err)
	}
	if cfg.shuffle {
		shuffleProblems(problems)
	}
	cnt := runQuiz(problems, cfg.limit)
	perc := 0.0
	if len(problems) > 0 {
		perc = 100 * float64(cnt) / float64(len(problems))
	}

	fmt.Printf("You managed to solve %d questions out of %d, %.2f%%\n", cnt, len(problems), perc)
}

func shuffleProblems(problems []problem) {
	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})
}
