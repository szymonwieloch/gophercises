package blackjack

type GameOption func(opts *gameOptions)

type gameOptions struct {
	games           uint
	decks           uint
	blackjackPayout float64
}

func defaultOptions() gameOptions {
	return gameOptions{
		games:           2,
		decks:           3,
		blackjackPayout: 1.5,
	}
}

// Set the number of consecutive game played in on session
func WithGames(games uint) GameOption {
	return func(opts *gameOptions) {
		opts.games = games
	}
}

// Set the number of card decks used by the game
func WithDecks(decks uint) GameOption {
	return func(opts *gameOptions) {
		opts.decks = decks
	}
}

// Sets the multiplayer for the blackjack reward
func WithBlackjackPayout(payout float64) GameOption {
	return func(opts *gameOptions) {
		opts.blackjackPayout = payout
	}
}
