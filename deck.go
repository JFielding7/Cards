package main

import (
	"math/rand"
	"strings"
)

type Deck struct {
	cards [DeckSize]*Card
	top   int
}

func NewOrderedDeck() *Deck {
	deck := new(Deck)

	for i := 0; i < DeckSize; i++ {
		deck.cards[i] = NewCard(i)
	}

	return deck
}

func NewShuffledDeck() *Deck {
	deck := NewOrderedDeck()

	rand.Shuffle(len(deck.cards), func(i, j int) {
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	})

	return deck
}

func (deck *Deck) DrawCard() *Card {
	if deck.top >= len(deck.cards) {
		return nil
	}

	card := deck.cards[deck.top]
	deck.top++
	return card
}

func (deck *Deck) String() string {
	var stringBuilder strings.Builder

	for i := 0; i < DeckSize; i++ {
		stringBuilder.WriteString(deck.cards[i].String() + " ")
	}

	return stringBuilder.String()
}
