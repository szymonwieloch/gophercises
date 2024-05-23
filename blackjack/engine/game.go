package engine

import "github.com/szymonwieloch/gophercises/cards"

type GameState int

const (
	Initial GameState = iota
	PlayerTurn
	DealerGetsCard
	DealerTurn
	PlayerWon
	DealerWon
	Tie
)

type Game struct {
	State      GameState
	Player     Player
	PlayerHand Hand
	DealerHand Hand
	deck       cards.Deck
}

func NewGame(player Player) Game {
	deck := cards.NewDeck(cards.WithDecks(6))
	return Game{
		Player:     player,
		PlayerHand: Hand{deck.Pop(), deck.Pop()},
		DealerHand: Hand{deck.Pop()},
		deck:       deck,
	}
}

func (game *Game) NextMove() bool {
	switch game.State {
	case Initial:
		playerHand, _ := HandValue(game.PlayerHand)
		if playerHand == BEST_HAND {
			game.State = PlayerWon
		} else {
			game.State = PlayerTurn
		}
	case PlayerTurn:
		makePlayerMove(game)
	case DealerGetsCard:
		game.DealerHand = append(game.DealerHand, game.deck.Pop())
		game.State = DealerTurn

	case DealerTurn:
		makeDealerMove(game)

	case PlayerWon, DealerWon, Tie:
		break
	default:
		panic("Invalid game state")

	}
	switch game.State {
	case PlayerWon, DealerWon, Tie:
		return false
	default:
		return true
	}
}

func makePlayerMove(game *Game) {
	switch game.Player.MakeDecision(game.PlayerHand, game.DealerHand) {
	case Hit:
		game.PlayerHand = append(game.PlayerHand, game.deck.Pop())
		playerHand, _ := HandValue(game.PlayerHand)
		if playerHand > BEST_HAND {
			game.State = DealerWon
		}
	case Stand:
		game.State = DealerGetsCard
	}
}

func shouldDealerContinue(game *Game) bool {
	val, soft17 := HandValue(game.DealerHand)
	return (val < 17) || (val == 17 && soft17)
}

func makeDealerMove(game *Game) {
	if shouldDealerContinue(game) {
		game.DealerHand = append(game.DealerHand, game.deck.Pop())
		return
	}
	// compare

	dealer, _ := HandValue(game.DealerHand)
	if dealer > BEST_HAND {
		game.State = PlayerWon
		return
	}
	player, _ := HandValue(game.PlayerHand)
	if player > dealer {
		game.State = PlayerWon
	} else if player < dealer {
		game.State = DealerWon
	} else {
		game.State = Tie
	}
}
