package main

import (
	"fmt"
	"strings"

	"github.com/szymonwieloch/gophercises/blackjack_ai/blackjack"
)

func handString(hand blackjack.Hand) string {
	score, _ := blackjack.Score(hand)
	var parts []string
	for _, card := range hand {
		parts = append(parts, card.String())
	}
	return fmt.Sprintf("%2d | ", score) + strings.Join(parts, ", ")
}

func drawBoard(playerHand blackjack.Hand, visibleDealerHand blackjack.Hand) {
	fmt.Println("Player: ", handString(playerHand))
	fmt.Println("Dealer: ", handString(visibleDealerHand))
}

type ConsolePlayer struct {
	firstDecision bool
}

func (player *ConsolePlayer) Bet() blackjack.Cents {
	fmt.Println("=========================================")
	fmt.Println("How many cents do you want to bet?")
	var result blackjack.Cents
	fmt.Scanf("%d", &result)
	fmt.Println("You decided to bet: ", result)
	return result
}

func (player *ConsolePlayer) OnStart(playerHand blackjack.Hand, visibleDealerHand blackjack.Hand) {
	drawBoard(playerHand, visibleDealerHand)
	player.firstDecision = true
}

func (player *ConsolePlayer) MakeDecision(playerHand blackjack.Hand, visibleDealerHand blackjack.Hand) blackjack.PlayerDecision {
	if !player.firstDecision {
		drawBoard(playerHand, visibleDealerHand)
	}
	player.firstDecision = false
	canSplit := blackjack.CanSplit(playerHand)
	splitStr := ""
	if canSplit {
		splitStr = ", s(P)lit"
	}

	for {
		fmt.Println("(H)it, (S)tand", splitStr, "or (D)ouble?")
		var decision string
		fmt.Scanln(&decision)
		clean := strings.ToLower(strings.TrimSpace(decision))
		switch clean {
		case "h", "hit":
			fmt.Println("Player decision: Hit")
			return blackjack.Hit
		case "s", "stand":
			fmt.Println("Player decision: Stand")
			return blackjack.Stand
		case "d", "double":
			fmt.Println("Player decision: Double")
			return blackjack.Double
		case "p", "split":
			if !canSplit {
				fmt.Println("Cannot split with this hand")
				break
			}
			return blackjack.Split
		default:
			fmt.Println("Unrecognised option: ", decision)
		}
	}
}

func (player *ConsolePlayer) OnGameCompleted(playerHand blackjack.Hand, dealerHand blackjack.Hand, gain blackjack.Cents) {
	drawBoard(playerHand, dealerHand)
	if gain > 0 {
		fmt.Println("You won ", gain)
	} else if gain == 0 {
		fmt.Println("Tie!")
	} else {
		fmt.Println("You lost ", -gain)
	}
}
