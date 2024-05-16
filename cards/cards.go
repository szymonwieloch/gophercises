package cards

import (
	"fmt"
	"math/rand"
	"slices"
)

//go:generate stringer -type=Suit,Rank

// Defines the card rank
type Rank uint8

// Defines the card suit
type Suit uint8

const (
	Clubs Suit = iota
	Diamonds
	Spades
	Hearts
	Joker
)

var allSuits = [...]Suit{Clubs, Diamonds, Spades, Hearts}

const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

var allRanks = [...]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

func rankShortString(r Rank) string {
	if r >= Two && r < Ten {
		return fmt.Sprint(r)
	}
	switch r {
	case Ace:
		return "A"
	case Ten:
		return "X"
	case King:
		return "K"
	case Queen:
		return "Q"
	case Jack:
		return "J"
	default:
		return ""
	}
}

func suitShortString(s Suit) string {
	switch s {
	case Hearts:
		return "♡"
	case Diamonds:
		return "♢"
	case Spades:
		return "♠"
	case Clubs:
		return "♣"
	default:
		return ""
	}
}

// Data type that describes a card
type Card struct {
	Rank Rank
	Suit Suit
}

func (card Card) String() string {
	if card.Suit == Joker {
		return "🃏"
	}
	return rankShortString(card.Rank) + suitShortString(card.Suit)
}

func (card Card) LongName() string {
	return fmt.Sprintln(card.Rank, "of", card.Suit)
}

// Collection of cards
type Deck []Card

// Shuffles the deck
func (deck Deck) Shuffle() {
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
}

// Sort
func (deck Deck) Sort() {
	slices.SortFunc(deck, func(a, b Card) int {
		if a.Suit >= b.Suit {
			return 1
		} else if a.Suit < b.Suit {
			return -1
		}
		return int(a.Rank) - int(b.Rank)
	})
}

type Option func(op *options)

type options struct {
	shuffle bool
	filter  func(c Card) bool
	jokers  uint
	decks   uint
}

// Returns default options
func newOptions() options {
	suits := make([]Suit, len(allSuits))
	copy(suits, allSuits[:])
	return options{
		shuffle: true,
		jokers:  0,
		decks:   1,
		filter:  func(c Card) bool { return false },
	}
}

func WithNoShuffle(o *options) {
	o.shuffle = false
}

func WithJokers(j uint) Option {
	return func(o *options) {
		o.jokers = j
	}
}

func WithDecks(d uint) Option {
	return func(o *options) {
		o.decks = d
	}
}

func WithFilter(filter func(c Card) bool) Option {
	return func(o *options) {
		o.filter = filter
	}
}

func NewDeck(opts ...Option) Deck {
	o := newOptions()
	for _, opt := range opts {
		opt(&o)
	}

	deck := Deck{}
	for range o.decks {
		for _, suit := range allSuits {
			for _, rank := range allRanks {
				card := Card{Suit: suit, Rank: rank}
				if o.filter(card) {
					continue
				}
				deck = append(deck, card)
			}
		}
	}

	for range o.jokers {
		deck = append(deck, Card{Suit: Joker})
	}
	if o.shuffle {
		deck.Shuffle()
	}
	return deck
}
