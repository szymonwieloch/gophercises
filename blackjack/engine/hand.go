package engine

import "github.com/szymonwieloch/gophercises/cards"

const BEST_HAND = 21

type Hand []cards.Card

// Returns the hand value (including making the best decision on Aces) and indication on the "soft 17"
func HandValue(hand Hand) (uint, bool) {
	var result uint = 0
	var aces int = 0
	for _, card := range hand {
		if card.Rank >= cards.Two && card.Rank <= cards.Ten {
			result += uint(card.Rank)
		} else if card.Rank == cards.Jack || card.Rank == cards.Queen || card.Rank == cards.King {
			result += 10
		} else {
			aces++
		}
	}
	if aces == 0 {
		return result, false
	}
	var elevens int
	for elevens = aces; elevens > 0; elevens-- {
		ones := aces - elevens
		if result+uint(11*elevens+1*ones) <= BEST_HAND {
			break
		}
	}
	result += uint(elevens*11) + uint(aces-elevens)
	if result == 17 {
		return result, elevens > 0
	} else {
		return result, false
	}

}
