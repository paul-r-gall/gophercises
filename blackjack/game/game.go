package game

import (
	"fmt"

	"github.com/paul-r-gall/gophercises/deck"
)

type Game struct {
	Players []*Player
	Deck    deck.Deck
	DSize   int
}

func NewGame(numPlayers int, numDecks int) *Game {
	g := Game{Players: []*Player{}, Deck: deck.NewDeck(nil, deck.MultDecks(numDecks)), DSize: numDecks * 52}
	g.Deck.Shuffle()

	for i := 0; i < numPlayers; i++ {
		g.Players = append(g.Players, NewPlayerCmd(false))
	}
	// dealer should be the last player.
	g.Players = append(g.Players, NewPlayerCmd(true))

	return &g
}

func (g *Game) Play() {
	numDrawn := 0
	end := false

	for !end {

		if numDrawn > g.DSize*4/5 {
			g.Deck.Reset()
			g.Deck.Shuffle()
			numDrawn = 0
		}
		// Place Bets
		for _, p := range g.Players {
			if !p.IsDealer {
				p.GetBet()
			}
		}

		// Deal Cards
		for _, p := range g.Players {
			p.Hit(g.Deck.Draw())
			numDrawn++
		}
		for i, p := range g.Players {
			p.Hit(g.Deck.Draw())
			numDrawn++
			if !p.IsDealer {
				fmt.Printf("Player %d: ", i)
				fmt.Println(p.CardsString())
			} else {
				fmt.Print("Dealer: ")
				fmt.Println(p.Cards[0].String() + ", |***|")
			}
		}

		// Hit/Stand
		for _, p := range g.Players {
			if p.IsDealer {
				fmt.Println(p.CardsString())
			}
			for {
				//check end turn conditions
				if p.IsBust() {
					fmt.Println("BUSTED")
					// they lose all their money
					p.BetAmt = 0
					break
				}
				if p.Has21() {
					fmt.Println("21!")
					break
				}

				// prompt player for input
				resp := p.GetResponse()

				if resp == Hit {
					p.Hit(g.Deck.Draw())
					numDrawn++
				} else if resp == Stand {
					break
				}
			}
		}

		// Resolve Bets
		dealer := g.Players[len(g.Players)-1]
		fmt.Printf("Dealer has "+dealer.CardsString()+": %d Points.\n", dealer.BestPoints())
		if dealer.IsBust() {
			for _, p := range g.Players {
				if p.IsDealer {
					continue
				}
				if p.Has21() {
					winPrint(p.BetAmt * 3 / 2)
					p.Money += (p.BetAmt * 5 / 2)
					p.BetAmt = 0

				} else {
					winPrint(p.BetAmt)
					p.Money += (p.BetAmt * 2)
					p.BetAmt = 0
				}
			}
		} else {
			for _, p := range g.Players {
				switch {
				case p.IsDealer:
					continue
				case p.IsBust():
					continue
				case dealer.Has21() && p.Has21():
					fmt.Println("Tie with Dealer")
					p.Money += p.BetAmt
					p.BetAmt = 0
				case !dealer.Has21() && p.Has21():
					winPrint(p.BetAmt * 3 / 2)
					p.Money += (p.BetAmt * 5 / 2)
					p.BetAmt = 0
				case p.BestPoints() > dealer.BestPoints():
					winPrint(p.BetAmt)
					p.Money += p.BetAmt * 2
					p.BetAmt = 0
				case p.BestPoints() == dealer.BestPoints():
					fmt.Println("Tie with Dealer")
					p.Money += p.BetAmt
					p.BetAmt = 0
				case p.BestPoints() < dealer.BestPoints():
					fmt.Println("You lose your bet.")
					p.BetAmt = 0
				}
			}
		}

		for _, p := range g.Players {
			p.Cards = []deck.Card{}
		}

		fmt.Println("Type Y to play again")
		r := ""
		fmt.Scanln(&r)
		if r != "Y" {
			end = true
		}
	}
}

func winPrint(amt int) {
	fmt.Printf("You Win $%d.\n", amt)
}
