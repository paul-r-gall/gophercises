package deck

import (
	"math/rand"
	"sort"
	"time"
)

// Deck is the user-safe exported deck
type Deck struct {
	cards []Card
	less  func(Card, Card) bool
}

//Shuffle does a random shuffle of the Deck.
func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := len(d.cards)
	sCards := make([]Card, l)
	perm := r.Perm(l)
	for i, num := range perm {
		sCards[i] = d.cards[num]
	}
	d.cards = sCards
}

//Sort sorts the deck according to the order defined by the user when the deck was created.
func (d *Deck) Sort() {
	sort.Slice(d.cards, func(i, j int) bool {
		return d.less(d.cards[i], d.cards[j])
	})
}

// TopCard returns the current top card of the deck (0-index of the internal array)
func (d Deck) TopCard() Card {
	return d.cards[0]
}

func defCards() []Card {
	var cards []Card
	for val := 1; val < 14; val++ {
		for suit := CLUBS; suit < 5; suit++ {
			cards = append(cards, Card{suit: suit, val: val})
		}
	}
	return cards
}

// NewDeck generates a new deck of cards using a user defined
// ordering and functional options for which cards to include.
// By default, if run as NewDeck(nil), it will use the standard
// ordering of 2 low, Ace high, and Clubs<Diamonds<Hearts<Spades
func NewDeck(less func(Card, Card) bool, opts ...func(*[]Card)) Deck {
	nd := Deck{}

	if less == nil {
		nd.less = defLess
	} else {
		nd.less = less
	}
	nd.cards = defCards()
	for _, fn := range opts {
		fn(&nd.cards)
	}
	return nd
}
