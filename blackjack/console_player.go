package main

import (
	"fmt"
	"strings"

	"github.com/szymonwieloch/gophercises/blackjack/engine"
)

type consolePlayer struct {
}

func (cp consolePlayer) MakeDecision(playerHand engine.Hand, dealerHand engine.Hand) engine.PlayerDecision {
	for {
		fmt.Println(("Make decision: Hit or Stand:"))
		var decision string
		fmt.Scanln(&decision)
		decision = strings.ToLower(decision)
		switch decision {
		case "h", "hit":
			fmt.Println("Player decision: Hit")
			return engine.Hit
		case "s", "stand":
			fmt.Println("Player decision: Stand")
			return engine.Stand
		default:
			fmt.Println("Unrecognised option: ", decision)
		}
	}
}
