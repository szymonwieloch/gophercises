package blackjack

import (
	"github.com/szymonwieloch/gophercises/cards"
)

// Plays a series of games, returns total gain
func Play(player Player, options ...GameOption) Cents {
	opts := applyOptions(options)

	var result Cents
	var deck cards.Deck
	for range opts.games {
		maybeResetDeck(&deck, opts.decks)
		result += playSingle(&deck, player, opts)
	}
	return result
}

func maybeResetDeck(deck *cards.Deck, deckCfg uint) {
	lenDeck := len(*deck)
	threshold := 52 * int(deckCfg) / 3
	if lenDeck < threshold {
		*deck = cards.NewDeck(cards.WithDecks(deckCfg))
	}
}

type handState struct {
	hand Hand
	bet  Cents
}

func playSingle(deck *cards.Deck, player Player, opts gameOptions) Cents {
	hs := handState{
		bet:  player.Bet(*deck),
		hand: Hand{deck.Pop(), deck.Pop()},
	}

	dealerHand := Hand{deck.Pop(), deck.Pop()}

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
	var totalResult Cents
	leftHands := []handState{hs}
	var processedHands []handState
	for len(leftHands) > 0 {
		current := leftHands[0]
		leftHands = leftHands[1:]
		current, newHands := playerTurn(current, player, dealerHand, deck)
		if isBusted(current.hand) {
			player.OnGameCompleted(current.hand, dealerHand, -current.bet)
			totalResult -= current.bet
		} else {
			processedHands = append(processedHands, current)
		}
		leftHands = append(leftHands, newHands...)
	}
	if len(processedHands) == 0 {
		return totalResult
	}

	for dealerNeedsToBid(dealerHand) {
		dealerHand = append(dealerHand, deck.Pop())
	}
	for _, ps := range processedHands {
		gameResult := checkResult(ps, dealerHand)
		player.OnGameCompleted(hs.hand, dealerHand, gameResult)
		totalResult += gameResult
	}
	return totalResult
}

// applies user decision. Returns indication if the hand is done, the hand state and new hand (if created by the split)
func applyPlayerDecision(hs handState, deck *cards.Deck, decision PlayerDecision) (bool, handState, Hand) {
	switch decision {
	case Stand:
		return true, hs, nil
	case Hit:
		hs.hand = append(hs.hand, deck.Pop())
		return isBusted(hs.hand), hs, nil
	case Double:
		hs.bet *= 2
		hs.hand = append(hs.hand, deck.Pop())
		return true, hs, nil
	case Split:
		if !CanSplit(hs.hand) {
			panic("Invalid user move: split")
		}
		newHand := hs.hand[1:2]
		hs.hand = hs.hand[:1]
		return false, hs, newHand
	default:
		panic("Unrechable")
	}
}

func playerTurn(hs handState, player Player, dealerHand Hand, deck *cards.Deck) (handState, []handState) {
	// return on stand or busted
	newHands := []handState{}
	for {
		decision := player.MakeDecision(hs.hand, dealerHand[:1])
		var newHand Hand
		var done bool
		done, hs, newHand = applyPlayerDecision(hs, deck, decision)
		if newHand != nil {
			newHands = append(newHands, handState{hand: newHand, bet: player.Bet(*deck)})
		}
		if done {
			return hs, newHands
		}
	}
}

// After the dealer collected all cards, performs check and side-effects on each hand state
func checkResult(hs handState, dealerHand Hand) Cents {
	if isBusted(dealerHand) {
		return hs.bet
	}
	dealerScore, _ := Score(dealerHand)
	playerScore, _ := Score(hs.hand)
	if dealerScore > playerScore {
		return -hs.bet
	} else if dealerScore < playerScore {
		return hs.bet
	} else {
		return 0
	}
}
