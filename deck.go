package main

import (
	"fmt"
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
	card := deck.cards[deck.top]
	deck.top = (deck.top + 1) % DeckSize
	return card
}

func (deck *Deck) String() string {
	var stringBuilder strings.Builder

	for i := 0; i < DeckSize; i++ {
		fmt.Println(deck.cards[i].value, deck.cards[i].suit)
		stringBuilder.WriteString(deck.cards[i].String() + " ")
	}

	return stringBuilder.String()
}
