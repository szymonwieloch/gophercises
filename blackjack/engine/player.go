package engine

type PlayerDecision int

const (
	Hit PlayerDecision = iota
	Stand
)

type Player interface {
	MakeDecision(playerHand Hand, dealerHand Hand) PlayerDecision
}
