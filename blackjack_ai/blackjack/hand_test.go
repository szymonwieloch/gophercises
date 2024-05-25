package blackjack

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		value, soft17 := Score(cs.hand)
		if value != cs.value {
			t.Error("Invalid value ", value, " for case ", cs.hand)
		}
		if soft17 != cs.soft17 {
			t.Error("Invalid soft 17 value ", soft17, " for case ", cs.hand)
		}
	}
}

func TestDealerNeedsToBid(t *testing.T) {
	cases := []struct {
		hand     Hand
		expected bool
	}{
		{
			hand:     Hand{cards.Card{Suit: cards.Clubs, Rank: cards.Ace}},
			expected: true,
		},
		{
			hand:     Hand{cards.Card{Suit: cards.Spades, Rank: cards.Ace}, cards.Card{Suit: cards.Clubs, Rank: cards.Ace}},
			expected: true,
		},
		{
			hand:     Hand{cards.Card{Suit: cards.Spades, Rank: cards.Ace}, cards.Card{Suit: cards.Clubs, Rank: cards.Six}},
			expected: true,
		},
		{
			hand:     Hand{cards.Card{Suit: cards.Spades, Rank: cards.Queen}, cards.Card{Suit: cards.Clubs, Rank: cards.Nine}},
			expected: false,
		},
	}

	for _, cs := range cases {
		if dealerNeedsToBid(cs.hand) != cs.expected {
			t.Error("Failed for case ", cs.hand)
		}
	}
}

func TestIsBusted(t *testing.T) {
	bustedHand := Hand{
		cards.Card{Rank: cards.Ten, Suit: cards.Spades},
		cards.Card{Rank: cards.Seven, Suit: cards.Spades},
		cards.Card{Rank: cards.Five, Suit: cards.Spades},
	}
	assert.True(t, isBusted(bustedHand))

	okHand := Hand{
		cards.Card{Rank: cards.Ten, Suit: cards.Spades},
		cards.Card{Rank: cards.Seven, Suit: cards.Spades},
	}
	assert.False(t, isBusted(okHand))
}

func TestIsBlackJack(t *testing.T) {
	bjHand := Hand{
		cards.Card{Rank: cards.Ten, Suit: cards.Spades},
		cards.Card{Rank: cards.Ace, Suit: cards.Spades},
	}
	assert.True(t, isBlackjack(bjHand))

	normalHand := Hand{
		cards.Card{Rank: cards.Seven, Suit: cards.Spades},
		cards.Card{Rank: cards.Ace, Suit: cards.Spades},
	}
	assert.False(t, isBlackjack(normalHand))

}
