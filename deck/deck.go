package deck

import (
	"math/rand"
	"sort"
	"time"
)

// Deck is the user-safe exported deck
type Deck struct {
	cards   []Card
	less    func(Card, Card) bool
	ogCards []Card
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

// Peek returns the current top card of the deck (0-index of the internal array)
func (d Deck) Peek() Card {
	return d.cards[0]
}

//Resets the deck to its original status -- replacing all drawn cards.
func (d *Deck) Reset() {
	d.cards = d.ogCards
}

// Draw is the same as Peek, but also removes the top card of the deck
func (d *Deck) Draw() Card {
	c := d.cards[0]
	d.cards = d.cards[1:]
	return c
}

// Here are some example functional options you can add.

// AddJokers returns a function which appends n Jokers to the end of the cards.
func AddJokers(n int) func(*[]Card) {
	return func(cs *[]Card) {
		for i := 0; i < n; i++ {
			*cs = append(*cs, Card{val: 0, suit: NONE})
		}
	}
}

// DeleteCards deletes the specific cards listed from the standard deck.
func DeleteCards(delCards map[Card]bool) func(*[]Card) {
	return func(cs *[]Card) {
		for i := 0; i < len(*cs); i++ {
			if delCards[(*cs)[i]] {
				*cs = append((*cs)[:i], (*cs)[i+1:]...)
				i--
			}
		}
	}
}

// MultDecks multiplies your deck by the provided number.
func MultDecks(numDecks int) func(*[]Card) {
	return func(cs *[]Card) {
		base := *cs
		for i := 1; i < numDecks; i++ {
			*cs = append(*cs, base...)
		}
	}
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
// These functional options are run in order, so be aware of that
// when using. The first option called will be run on the base
// 52-card deck
//
// By default, if run as NewDeck(nil), it will use the standard bridge
// ordering of Ace low, King high, and Clubs < Diamonds < Hearts < Spades
func NewDeck(less func(Card, Card) bool, opts ...func(*[]Card)) Deck {
	nd := Deck{}

	if less == nil {
		nd.less = defLess
	} else {
		nd.less = less
	}
	nd.ogCards = defCards()
	for _, fn := range opts {
		fn(&nd.ogCards)
	}
	nd.cards = nd.ogCards
	return nd
}
