package main

import "flag"

type config struct {
	ai              bool
	games           uint
	decks           uint
	blackjackPayout float64
}

func parseConfig() config {
	ai := flag.Bool("ai", false, "Enables AI instead of human")
	games := flag.Uint("games", 3, "Number of games to play")
	decks := flag.Uint("decks", 4, "Number of decks of cards to use")
	payout := flag.Float64("payout", 1.5, "Blackjack payout")
	flag.Parse()
	return config{
		ai:              *ai,
		games:           *games,
		decks:           *decks,
		blackjackPayout: *payout,
	}
}
