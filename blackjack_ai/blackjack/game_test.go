package blackjack

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/szymonwieloch/gophercises/cards"
)

func TestMaybeResetDeck(t *testing.T) {
	tests := []struct {
		deckLen     int
		deckCnt     uint
		expectedLen int
	}{
		{
			deckLen:     50,
			deckCnt:     1,
			expectedLen: 50,
		},
		{
			deckLen:     16,
			deckCnt:     1,
			expectedLen: 52,
		},
		{
			deckLen:     112,
			deckCnt:     3,
			expectedLen: 112,
		},
		{
			deckLen:     50,
			deckCnt:     3,
			expectedLen: 156,
		},
	}
	for _, test := range tests {
		deck := cards.NewDeck(cards.WithDecks(test.deckCnt))
		deck = deck[:test.deckLen]
		maybeResetDeck(&deck, test.deckCnt)
		assert.Equal(t, len(deck), test.expectedLen)
	}
}

func TestCheckResult(t *testing.T) {
	tests := []struct {
		player   handState
		dealer   Hand
		expected Cents
	}{
		{ // dealer busted
			player:   handState{hand: Hand{cards.Card{Suit: cards.Clubs, Rank: cards.Eight}}, bet: 100},
			dealer:   Hand{cards.Card{Suit: cards.Clubs, Rank: cards.Eight}, cards.Card{Suit: cards.Clubs, Rank: cards.Ten}, cards.Card{Suit: cards.Clubs, Rank: cards.Four}},
			expected: 100,
		},
		{ // player wins
			player:   handState{hand: Hand{cards.Card{Suit: cards.Clubs, Rank: cards.Eight}, cards.Card{Suit: cards.Clubs, Rank: cards.Ace}}, bet: 100},
			dealer:   Hand{cards.Card{Suit: cards.Clubs, Rank: cards.Eight}, cards.Card{Suit: cards.Clubs, Rank: cards.Ten}},
			expected: 100,
		},
		{ // player looses
			player:   handState{hand: Hand{cards.Card{Suit: cards.Clubs, Rank: cards.Eight}, cards.Card{Suit: cards.Clubs, Rank: cards.Nine}}, bet: 100},
			dealer:   Hand{cards.Card{Suit: cards.Clubs, Rank: cards.Eight}, cards.Card{Suit: cards.Clubs, Rank: cards.Ten}},
			expected: -100,
		},
		{ // player looses
			player:   handState{hand: Hand{cards.Card{Suit: cards.Clubs, Rank: cards.Nine}, cards.Card{Suit: cards.Clubs, Rank: cards.Nine}}, bet: 100},
			dealer:   Hand{cards.Card{Suit: cards.Clubs, Rank: cards.Eight}, cards.Card{Suit: cards.Clubs, Rank: cards.Ten}},
			expected: 0,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, checkResult(test.player, test.dealer))
	}
}

func TestApplyPlayerDecision(t *testing.T) {
	tests := []struct {
		hs              handState
		deck            cards.Deck
		decision        PlayerDecision
		expectedDone    bool
		expectedHs      handState
		expectedNewHand Hand
	}{
		{
			hs:           handState{hand: Hand{cards.Card{Rank: cards.Ten}}, bet: 100},
			decision:     Stand,
			expectedDone: true,
			expectedHs:   handState{hand: Hand{cards.Card{Rank: cards.Ten}}, bet: 100},
		}, {
			hs:           handState{hand: Hand{cards.Card{Rank: cards.Ten}}, bet: 100},
			decision:     Hit,
			deck:         cards.Deck{cards.Card{Rank: cards.Seven}},
			expectedDone: false,
			expectedHs:   handState{hand: Hand{cards.Card{Rank: cards.Ten}, cards.Card{Rank: cards.Seven}}, bet: 100},
		}, {
			hs:           handState{hand: Hand{cards.Card{Rank: cards.Ten}}, bet: 100},
			decision:     Double,
			deck:         cards.Deck{cards.Card{Rank: cards.Seven}},
			expectedDone: true,
			expectedHs:   handState{hand: Hand{cards.Card{Rank: cards.Ten}, cards.Card{Rank: cards.Seven}}, bet: 200},
		}, {
			hs:              handState{hand: Hand{cards.Card{Rank: cards.Ten, Suit: cards.Hearts}, cards.Card{Rank: cards.Ten, Suit: cards.Clubs}}, bet: 100},
			decision:        Split,
			expectedDone:    false,
			expectedHs:      handState{hand: Hand{cards.Card{Rank: cards.Ten, Suit: cards.Hearts}}, bet: 100},
			expectedNewHand: Hand{cards.Card{Rank: cards.Ten, Suit: cards.Clubs}},
		},
	}
	for _, test := range tests {
		done, hs, newHand := applyPlayerDecision(test.hs, &test.deck, test.decision)
		assert.Equal(t, done, test.expectedDone)
		assert.Equal(t, hs, test.expectedHs)
		assert.Equal(t, newHand, test.expectedNewHand)
	}

}
