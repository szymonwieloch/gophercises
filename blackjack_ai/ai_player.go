package main

import "github.com/szymonwieloch/gophercises/blackjack_ai/blackjack"

type AIPlayer struct {
}

func (player AIPlayer) Bet() blackjack.Cents {
	return 100
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
