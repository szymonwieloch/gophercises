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
	var deck cards.Deck = cards.NewDeck(cards.WithDecks(opts.decks))
	for range opts.games {
		// TODO reset deck
		result += playSingle(&deck, player, opts)
	}
	return result
}

func playSingle(deck *cards.Deck, player Player, opts gameOptions) Cents {
	bet := player.Bet()

	playerHand := Hand{deck.Pop(), deck.Pop()}
	dealerHand := Hand{deck.Pop(), deck.Pop()}

	player.OnStart(playerHand, dealerHand[:1])

	if isBlackjack(playerHand) && isBlackjack(dealerHand) {
		player.OnGameCompleted(playerHand, dealerHand, 0)
		return 0
	}
	if isBlackjack(dealerHand) {
		player.OnGameCompleted(playerHand, dealerHand, -bet)
		return -bet
	}
	if isBlackjack(playerHand) {
		result := Cents(float64(bet) * opts.blackjackPayout)
		player.OnGameCompleted(playerHand, dealerHand, result)
		return result
	}
loop:
	for {
		switch player.MakeDecision(playerHand, dealerHand[:1]) {
		case Stand:
			break loop
		case Hit:
			playerHand = append(playerHand, deck.Pop())
		}
		if isBusted(playerHand) {
			player.OnGameCompleted(playerHand, dealerHand, -bet)
			return -bet
		}
	}

	for dealerNeedsToBid(dealerHand) {
		dealerHand = append(dealerHand, deck.Pop())
	}
	if isBusted(dealerHand) {
		player.OnGameCompleted(playerHand, dealerHand, bet)
		return bet
	}
	dealerScore, _ := Score(dealerHand)
	playerScore, _ := Score(playerHand)
	if dealerScore > playerScore {
		player.OnGameCompleted(playerHand, dealerHand, -bet)
		return -bet
	} else if dealerScore < playerScore {
		player.OnGameCompleted(playerHand, dealerHand, bet)
		return bet
	} else {
		player.OnGameCompleted(playerHand, dealerHand, 0)
		return 0
	}
}
