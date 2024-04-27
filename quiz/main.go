package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/szymonwieloch/gophercises/quiz/questions"
)

func main() {
	csvPath := flag.String("p", "problems.csv", "Path to the CSV with problems")
	flag.Parse()
	quests, err := questions.ParseFile(*csvPath)
	if err != nil {
		log.Fatal(err)
	}
	cnt, err := runQuiz(quests)
	if err != nil {
		log.Fatal(err)
	}
	perc := 0.0
	if len(quests) > 0 {
		perc = 100 * float64(cnt) / float64(len(quests))
	}

	fmt.Printf("You managed to solve %d questions out of %d, %.2f%%\n", cnt, len(quests), perc)
}

// Shows user the question return indicates if the user got it correctly
func showQuestion(q *questions.Question) (bool, error) {
	fmt.Print(q.Question, "?: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	text = strings.TrimSpace(text)
	return (text == q.Answer), nil
}

// Return the number of questions user got right
func runQuiz(quests []questions.Question) (int, error) {
	result := 0
	for _, q := range quests {
		ok, err := showQuestion(&q)
		if err != nil {
			return 0, err
		}
		if ok {
			result++
		}
	}
	return result, nil
}
