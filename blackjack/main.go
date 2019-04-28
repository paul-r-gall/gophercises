package main

import (
	"fmt"

	"github.com/paul-r-gall/gophercises/blackjack/game"
)

func main() {
	fmt.Println("How many players?")
	numPlayers := 1
	fmt.Scanln(&numPlayers)
	g := game.NewGame(numPlayers, 4)
	g.Play()
}
