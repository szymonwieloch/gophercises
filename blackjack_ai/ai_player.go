package main

import (
	"github.com/szymonwieloch/gophercises/blackjack_ai/blackjack"
	"github.com/szymonwieloch/gophercises/cards"
)

func deckScore(deck cards.Deck) int {
	result := 0
	for idx := range deck {
		score, _ := blackjack.Score(blackjack.Hand(deck[idx : idx+1]))
		if score >= 10 {
			result++
		}
		if score <= 6 {
			result--
		}
	}
	return result
}

func leftDecks(deck cards.Deck) int {
	return len(deck) / 52
}

type AIPlayer struct {
}

func (player AIPlayer) Bet(deck cards.Deck) blackjack.Cents {
	score := deckScore(deck)
	trueScore := score / leftDecks(deck)
	// if trueScore > 10 {
	// 	fmt.Println(trueScore)
	// }
	switch {
	case trueScore > 14:
		return 10000
	case trueScore > 8:
		return 500
	default:
		return 100
	}
}

func (player AIPlayer) MakeDecision(playerHand blackjack.Hand, visibleDealerHand blackjack.Hand) blackjack.PlayerDecision {
	score, soft := blackjack.Score(playerHand)
	if blackjack.CanSplit(playerHand) {
		cardScore, _ := blackjack.Score(playerHand[:1])
		if cardScore >= 8 && cardScore != 10 {
			return blackjack.Split
		}
	}
	if (score == 10 || score == 11) && !soft {
		return blackjack.Double
	}
	dealerScore, _ := blackjack.Score(visibleDealerHand)
	if dealerScore >= 5 && dealerScore <= 6 {
		return blackjack.Stand
	}
	if score < 13 {
		return blackjack.Hit
	} else {
		return blackjack.Stand
	}
}

func (player AIPlayer) OnGameCompleted(playerHand blackjack.Hand, dealerHand blackjack.Hand, gain blackjack.Cents) {

}

var _ blackjack.Player = (*AIPlayer)(nil)
