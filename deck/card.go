package deck

import (
	"strconv"
)

// Suit ranges from 0-4, possibilities are
// NONE (joker), CLUBS, DIAMOND, HEARTS, SPADES
type Suit int

//
const (
	NONE Suit = iota
	CLUBS
	DIAMONDS
	HEARTS
	SPADES
)

// Card -- Suit ranges from Clubs to Spades, Val from 1-13 (1 being Ace, 13 being King)
// A Suit of NONE implies the card is a Joker. The Val of a Joker should be zero.
// These internal variables are not changeable from outside the package.
type Card struct {
	suit Suit
	val  int
}

// Suit returns the suit of the card as a Suit
func (c Card) Suit() Suit {
	return c.suit
}

func defLess(c1 Card, c2 Card) bool {
	if c1.suit < c2.suit {
		return true
	}
	if c1.val == 1 {
		return false
	}
	if c2.val == 1 {
		return true
	}
	if c2.val > c1.val {
		return true
	}
	return false
}

// Denom returns the Denomination of the card as an int
func (c Card) Denom() int {
	return c.val
}

// Name gives the name of the card as a string.
func (c Card) Name() string {
	if c.suit == NONE {
		return "Joker"
	}
	s := ""
	switch c.val {
	case 1:
		s += "Ace"
	case 11:
		s += "Jack"
	case 12:
		s += "Queen"
	case 13:
		s += "King"
	default:
		s += strconv.Itoa(c.val)
	}
	s += " of "
	switch c.suit {
	case CLUBS:
		s += "Clubs"
	case HEARTS:
		s += "Hearts"
	case SPADES:
		s += "Spades"
	case DIAMONDS:
		s += "Diamonds"
	}
	return s
}
