package blackjack

import (
	"github.com/szymonwieloch/gophercises/cards"
)

// Plays a series of games, returns total gain
func Play(player Player, options ...GameOption) Cents {
	opts := defaultOptions()
	for _, opt := range options {
		opt(&opts)
	}
	var result Cents
	var deck cards.Deck
	for range opts.games {
		maybeResetDeck(&deck, opts.decks)
		result += playSingle(&deck, player, opts)
	}
	return result
}

func maybeResetDeck(deck *cards.Deck, deckCfg uint) {
	if len(*deck) < (52 * int(deckCfg) / 3) {
		*deck = cards.NewDeck(cards.WithDecks(deckCfg))
	}
}

type handState struct {
	hand Hand
	bet  Cents
}

func playSingle(deck *cards.Deck, player Player, opts gameOptions) Cents {
	hs := handState{
		bet:  player.Bet(),
		hand: Hand{deck.Pop(), deck.Pop()},
	}

	dealerHand := Hand{deck.Pop(), deck.Pop()}

	player.OnStart(hs.hand, dealerHand[:1])

	if isBlackjack(hs.hand) && isBlackjack(dealerHand) {
		player.OnGameCompleted(hs.hand, dealerHand, 0)
		return 0
	}
	if isBlackjack(dealerHand) {
		player.OnGameCompleted(hs.hand, dealerHand, -hs.bet)
		return -hs.bet
	}
	if isBlackjack(hs.hand) {
		result := Cents(float64(hs.bet) * opts.blackjackPayout)
		player.OnGameCompleted(hs.hand, dealerHand, result)
		return result
	}
	playerTurn(&hs, player, dealerHand, deck)
	if isBusted(hs.hand) {
		player.OnGameCompleted(hs.hand, dealerHand, -hs.bet)
		return -hs.bet
	}

	for dealerNeedsToBid(dealerHand) {
		dealerHand = append(dealerHand, deck.Pop())
	}
	return checkResult(hs, dealerHand, player)
}

func playerTurn(hs *handState, player Player, dealerHand Hand, deck *cards.Deck) []handState {
	// return on stand or busted
	newHands := []handState{}
	for {
		switch player.MakeDecision(hs.hand, dealerHand[:1]) {
		case Stand:
			return newHands
		case Hit:
			hs.hand = append(hs.hand, deck.Pop())
		case Double:
			hs.bet *= 2
			hs.hand = append(hs.hand, deck.Pop())
			return newHands
		case Split:
			if !CanSplit(hs.hand) {
				panic("Invalid user move: split")
			}
			hs.hand = hs.hand[:1]
			newHands = append(newHands, handState{
				hand: hs.hand[1:],
				bet:  player.Bet(),
			})
			// newHand :=
		}
		if isBusted(hs.hand) {
			return newHands
		}
	}
}

// After the dealer collected all cards, performs check and side-effects on each hand state
func checkResult(hs handState, dealerHand Hand, player Player) Cents {
	if isBusted(dealerHand) {
		player.OnGameCompleted(hs.hand, dealerHand, hs.bet)
		return hs.bet
	}
	dealerScore, _ := Score(dealerHand)
	playerScore, _ := Score(hs.hand)
	if dealerScore > playerScore {
		player.OnGameCompleted(hs.hand, dealerHand, -hs.bet)
		return -hs.bet
	} else if dealerScore < playerScore {
		player.OnGameCompleted(hs.hand, dealerHand, hs.bet)
		return hs.bet
	} else {
		player.OnGameCompleted(hs.hand, dealerHand, 0)
		return 0
	}
}
