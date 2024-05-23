package main

import (
	"fmt"
	"strings"

	"github.com/szymonwieloch/gophercises/blackjack/engine"
)

func handString(hand engine.Hand) string {
	value, _ := engine.HandValue(hand)
	var parts []string
	for _, card := range hand {
		parts = append(parts, card.String())
	}
	return fmt.Sprintf("%2d | ", value) + strings.Join(parts, ", ")
}

func printGameState(game *engine.Game) {
	fmt.Println("--------------------------------------")
	fmt.Println()
	fmt.Println("Dealer:", handString(game.DealerHand))
	fmt.Println("Player:", handString(game.PlayerHand))
	fmt.Println()
}

func main() {
	fmt.Println("Blackjack!")
	fmt.Println()
	game := engine.NewGame(consolePlayer{})
	for game.NextMove() {
		printGameState(&game)
	}
	printGameState(&game)
	switch game.State {
	case engine.PlayerWon:
		fmt.Println("You won!")
	case engine.DealerWon:
		fmt.Println("Dealer won!")
	case engine.Tie:
		fmt.Println("Tie!")
	}

}
