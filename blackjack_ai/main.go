package main

import (
	"fmt"

	"github.com/szymonwieloch/gophercises/blackjack_ai/blackjack"
)

func main() {
	fmt.Println("Blackjack!")
	cfg := parseConfig()
	var player blackjack.Player
	if cfg.ai {
		player = AIPlayer{}
	} else {
		player = &ConsolePlayer{}
	}
	gained := blackjack.Play(player, blackjack.WithBlackjackPayout(cfg.blackjackPayout), blackjack.WithDecks(cfg.decks), blackjack.WithGames(cfg.games))
	result := "won"
	if gained < 0 {
		result = "lost"
		gained = -gained
	}
	fmt.Println("All games completed, you", result, "in total", gained)
}
