package blackjack

type PlayerDecision int

const (
	Hit PlayerDecision = iota
	Stand
	Double
	Split
)

type Player interface {
	Bet() Cents
	OnStart(playerHand Hand, visibleDealerHand Hand)
	MakeDecision(playerHand Hand, visibleDealerHand Hand) PlayerDecision
	OnGameCompleted(playerHand Hand, dealerHand Hand, gain Cents)
}
