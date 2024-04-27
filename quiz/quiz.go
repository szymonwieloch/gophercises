package main

import (
	"fmt"
	"strings"
	"time"
)

// Shows user the question return indicates if the user got it correctly
func showProblem(q *problem) bool {
	var answer string
	for {
		fmt.Print(q.question, "?: ")
		_, err := fmt.Scanf("%s\n", &answer)
		if err == nil {
			break
		}
		fmt.Println("Invalid answer: ", err.Error())
	}
	answer = strings.TrimSpace(answer)
	return (answer == q.answer)
}

// Rus the quiz in a coroutine, sends info if the answer was correct
func runBackgroudQuiz(problems []problem, results chan bool) {
	for _, problem := range problems {
		results <- showProblem(&problem)
	}
	close(results)
}

// Return the number of questions user got right
func runQuiz(problems []problem, limit time.Duration) int {
	result := 0
	results := make(chan bool)
	var timerChan <-chan time.Time
	if limit != 0 {
		timerChan = time.NewTimer(limit).C
	}
	go runBackgroudQuiz(problems, results)
forLabel:
	for {
		select {
		case res, ok := <-results:
			if !ok {
				break forLabel
			}
			if res {
				result++
			}
		case <-timerChan:
			fmt.Println("Time limit!")
			break forLabel

		}
	}
	return result
}
