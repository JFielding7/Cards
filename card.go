package main

import "fmt"

const (
	Values      = 13 // Number of values in a suit
	ValueOffset = 2  // Offset for card values
	DeckSize    = 52 // Total number of cards in a deck
)

const (
	TEN   = iota + 10
	Jack  // Jack starts at value 11
	Queen // Queen is 12
	King  // King is 13
	Ace   // Ace is 14
)

const (
	Clubs    = iota // Clubs suit
	Hearts          // Hearts suit
	Spades          // Spades suit
	Diamonds        // Diamonds suit
)

type Card struct {
	suit  int
	value int
}

func NewCard(index int) *Card {
	return &Card{suit: index / Values, value: index%Values + ValueOffset}
}

func (c *Card) ValueChar() rune {
	switch c.value {
	case TEN:
		return 'T'
	case Jack:
		return 'J'
	case Queen:
		return 'Q'
	case King:
		return 'K'
	case Ace:
		return 'A'
	default:
		return rune(c.value) + '0'
	}
}

func (c *Card) SuitChar() rune {
	switch c.suit {
	case Clubs:
		return '♣'
	case Hearts:
		return '♥'
	case Spades:
		return '♠'
	case Diamonds:
		return '♦'
	default:
		panic("Invalid Suit")
	}
}

func (c *Card) String() string {
	return fmt.Sprintf("%c%c", c.ValueChar(), c.SuitChar())
}
