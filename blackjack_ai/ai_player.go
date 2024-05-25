package main

import "github.com/szymonwieloch/gophercises/blackjack_ai/blackjack"

type AIPlayer struct {
}

func (player AIPlayer) Bet() blackjack.Cents {
	return 100
}

func (player AIPlayer) MakeDecision(playerHand blackjack.Hand, visibleDealerHand blackjack.Hand) blackjack.PlayerDecision {
	return blackjack.Hit
}
