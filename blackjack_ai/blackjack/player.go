package blackjack

import "github.com/szymonwieloch/gophercises/cards"

type PlayerDecision int

const (
	Hit PlayerDecision = iota
	Stand
	Double
	Split
)

type Player interface {
	Bet(deck cards.Deck) Cents
	MakeDecision(playerHand Hand, visibleDealerHand Hand) PlayerDecision
	OnGameCompleted(playerHand Hand, dealerHand Hand, gain Cents)
}
