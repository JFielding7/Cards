package main

import (
	"fmt"
	"strings"
)

const (
	WIN = iota
	DRAW
	LOSS
)

const (
	RoyalValue    = 10
	StartingCards = 2
	Target        = 21
	AceChange     = 10
)

type Hand struct {
	cards  []*Card
	sum    int
	hasAce bool
}

type HitStrategy func(*Hand) bool

func (hand *Hand) AddCard(card *Card) {
	hand.cards = append(hand.cards, card)

	switch card.value {
	case Ace:
		hand.sum++
		hand.hasAce = true
	case Jack, Queen, King:
		hand.sum += RoyalValue
	default:
		hand.sum += card.value
	}
}

func DealHands(deck *Deck, playerHand *Hand, opponentHand *Hand) {
	for i := 0; i < StartingCards; i++ {
		card := deck.DrawCard()
		playerHand.AddCard(card)

		card = deck.DrawCard()
		opponentHand.AddCard(card)
	}
}

func (hand *Hand) score() int {
	score := hand.sum
	if hand.hasAce && hand.sum <= Target-AceChange {
		score += AceChange
	}
	return score
}

func (hand *Hand) String() string {
	var stringBuilder strings.Builder
	for _, card := range hand.cards {
		stringBuilder.WriteString(card.String() + " ")
	}
	return stringBuilder.String()
}

func PrintHands(playerHand *Hand, opponentHand *Hand) {
	fmt.Printf("Player:   %d\n", playerHand.score())
	fmt.Println(playerHand)
	fmt.Printf("Opponent: %d\n", opponentHand.score())
	fmt.Println(opponentHand)
	fmt.Println()
}

func isWinner(player int, other int) bool {
	return player <= Target && (other > Target || player > other)
}

func PlayBlackjack(playerStrategy HitStrategy, opponentStrategy HitStrategy) int {
	deck := NewShuffledDeck()
	playerHand := new(Hand)
	opponentHand := new(Hand)

	DealHands(deck, playerHand, opponentHand)

	//PrintHands(playerHand, opponentHand)

	for {
		playerHits := playerStrategy(playerHand)
		opponentHits := opponentStrategy(opponentHand)

		if !playerHits && !opponentHits {
			break
		}
		if playerHits {
			playerHand.AddCard(deck.DrawCard())
		}
		if opponentHits {
			opponentHand.AddCard(deck.DrawCard())
		}

		//PrintHands(playerHand, opponentHand)
	}

	playerScore := playerHand.score()
	opponentScore := opponentHand.score()
	if isWinner(playerScore, opponentScore) {
		return WIN
	} else if isWinner(opponentScore, playerScore) {
		return LOSS
	}
	return DRAW
}
