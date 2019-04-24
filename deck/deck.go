package deck

// Deck is the user-safe exported deck
type Deck struct {
	cards []Card
	less  func(Card, Card) bool
}

//Shuffle does a random shuffle of the Deck.
func (d *Deck) Shuffle() {

}

//Sort sorts the deck according to the order defined by the user when the deck was created.
func (d *Deck) Sort() {

}

// TopCard returns the current top card of the deck (0-index of the internal array)
func (d Deck) TopCard() *Card {
	return nil
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
func NewDeck(less func(Card, Card) bool, opts ...func(*[]Card)) *Deck {
	nd := Deck{}

	if less == nil {
		nd.less = defLess
	} else {
		nd.less = less
	}
	return nil
}
