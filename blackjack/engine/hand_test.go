package engine

import (
	"testing"

	"github.com/szymonwieloch/gophercises/cards"
)

func TestHandValue(t *testing.T) {
	cases := []struct {
		hand   Hand
		value  uint
		soft17 bool
	}{
		{
			hand:   Hand{cards.Card{Suit: cards.Clubs, Rank: cards.Eight}},
			value:  8,
			soft17: false,
		},
		{
			hand:   Hand{cards.Card{Suit: cards.Clubs, Rank: cards.Two}, cards.Card{Suit: cards.Hearts, Rank: cards.Queen}},
			value:  12,
			soft17: false,
		},
		{
			hand:   Hand{cards.Card{Suit: cards.Clubs, Rank: cards.Ace}},
			value:  11,
			soft17: false,
		},
		{
			hand:   Hand{cards.Card{Suit: cards.Spades, Rank: cards.Ace}, cards.Card{Suit: cards.Clubs, Rank: cards.Ace}},
			value:  12,
			soft17: false,
		},
		{
			hand:   Hand{cards.Card{Suit: cards.Spades, Rank: cards.Ace}, cards.Card{Suit: cards.Clubs, Rank: cards.Six}},
			value:  17,
			soft17: true,
		},
		{
			hand:   Hand{cards.Card{Suit: cards.Spades, Rank: cards.King}, cards.Card{Suit: cards.Clubs, Rank: cards.Seven}},
			value:  17,
			soft17: false,
		},
	}
	for _, cs := range cases {
		value, soft17 := HandValue(cs.hand)
		if value != cs.value {
			t.Error("Invalid value ", value, " for case ", cs.hand)
		}
		if soft17 != cs.soft17 {
			t.Error("Invalid soft 17 value ", soft17, " for case ", cs.hand)
		}
	}
}
