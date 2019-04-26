package game

import (
	deck "github.com/paul-r-gall/gophercises/deck"
)

type Response int

const (
	Hit Response = iota
	Stand
)

type Player struct {
	IsDealer bool
	Money    int
	Cards    []deck.Card
	BetAmt   int
	
	// Idea: these two should be customized based on the
	// hosting of the game.
	responseMethod func() Response
	betMethod      func() int

	// Channels for input and output. Again, we are aiming
	// to expand to a webhosted game.
	RespInChan  <-chan Response
	RespOutChan chan<- Response

	BetInChan  <-chan int
	BetOutChan chan<- int

	GenOutChan chan<- string
}

func NewPlayer(isDealer bool, cmd bool) *Player {
	p := Player{}
	if isDealer {
		p.IsDealer = isDealer
		p.responseMethod = func() Response {
			pts := p.Points()
			// Start by counting A=11: must Hit if lower than 17, must stand otherwise
			if pts[1] < 17 {
				return Hit
			}
			if pts[1] < 22 {
				return Stand
			}
			// Counting A=1: must Hit if less than 17, must stand otherwise
			if pts[0] < 17 {
				return Hit
			}
			return Stand
		}
		p.betMethod = nil

	} else if cmd {
		p.Money = 500
		p.IsDealer = isDealer
		p.responseMethod = func() Response {
			stand := false
			for !stand && !p.IsBust() {

			}

		}

	}

	return &p
}

func NewPlayerCmd(isDealer bool) *Player {
	return NewPlayer(isDealer, true)
}

// Points returns an ordered pair. The first number
// is the score if A=1. The second number is the score
// if A=11.
func (p Player) Points() [2]int {
	s1 := 0
	s2 := 0
	for _, c := range p.Cards {
		switch {
		case c.Denom() == 1:
			s1 += 1
			s2 += 11
		case c.Denom() > 9:
			s1 += 10
			s2 += 10
		default:
			s1 += c.Denom()
			s2 += c.Denom()
		}
	}
	return [2]int{s1, s2}
}

// IsBust returns whether the player has gone bust.
func (p Player) IsBust() bool {
	if p.Points()[0] > 21 {
		return true
	}
	return false
}

func (p *Player) Hit(c deck.Card) {
	p.Cards = append(p.Cards, c)
}

func (p Player) GetResponse() Response {
	return p.responseMethod()
}

func (p Player) GetBet() int {
	return p.betMethod()
}
