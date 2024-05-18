package cards

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleCard() {
	fmt.Println(Card{Suit: Hearts, Rank: Ten})
	fmt.Println(Card{Suit: Clubs, Rank: King}.LongName())

	//Output:
	// 10♥
	// King of Clubs
}

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	assert.Equal(t, len(deck), 52)
}

func TestAddJokers(t *testing.T) {
	deck := NewDeck(WithJokers(3))
	assert.Equal(t, len(deck), 55)
}

func TestRemoveRank(t *testing.T) {
	deck := NewDeck(WithFilter(func(c Card) bool {
		return c.Rank == Two || c.Rank == Three
	}))
	assert.Equal(t, len(deck), 44)
	assert.False(t, slices.Contains(deck, Card{Rank: Three, Suit: Spades}))
}

func TestRemoveSuit(t *testing.T) {
	deck := NewDeck(WithFilter(func(c Card) bool {
		return c.Suit == Spades
	}))
	assert.Equal(t, len(deck), 39)
	assert.False(t, slices.Contains(deck, Card{Rank: Five, Suit: Spades}))
}

func TestShuffle(t *testing.T) {
	sorted1 := NewDeck(WithNoShuffle)
	sorted2 := NewDeck(WithNoShuffle)
	random := NewDeck()

	assert.Equal(t, sorted1, sorted2)
	assert.NotEqual(t, sorted1, random)

}

func TestDeckNumber(t *testing.T) {
	deck := NewDeck(WithDecks(3))
	assert.Equal(t, len(deck), 3*52)
}
