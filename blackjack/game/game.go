package game

import (
	"github.com/paul-r-gall/gophercises/deck"
)

type Game struct {
	Players []Player
	Deck    deck.Deck
}

func NewGame(numPlayers int, numDecks int) *Game {
	return nil
}
