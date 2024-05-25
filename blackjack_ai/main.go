package main

import (
	"fmt"

	"github.com/szymonwieloch/gophercises/blackjack_ai/blackjack"
)

func main() {
	fmt.Println("Blackjack!")
	var player ConsolePlayer
	gained := blackjack.Play(&player)
	result := "won"
	if gained < 0 {
		result = "lost"
		gained = -gained
	}
	fmt.Println("All games completed, you", result, "in total", gained)
}
