package blackjack

import "github.com/szymonwieloch/gophercises/cards"

const BestScore = 21

type Hand []cards.Card

// Returns the hand score (including making the best decision on Aces) and indication on the "soft 17"
func Score(hand Hand) (uint, bool) {
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
		if result+uint(11*elevens+1*ones) <= BestScore {
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

// Dealer logic: check if dealer needs to continue bidding
func dealerNeedsToBid(hand Hand) bool {
	score, soft := Score(hand)
	return score < 17 || (score == 17 && soft)
}

func isBlackjack(hand Hand) bool {
	if len(hand) != 2 {
		return false
	}
	score, _ := Score(hand)
	return score == BestScore
}

func isBusted(hand Hand) bool {
	score, _ := Score(hand)
	return score > BestScore
}

func CanSplit(hand Hand) bool {
	return len(hand) == 2 && hand[0].Rank == hand[1].Rank
}
