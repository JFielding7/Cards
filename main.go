package main

import "fmt"

const (
	Rounds = 1 << 25
)

func PlayerHitStrategy(hand *Hand) bool {
	//return !(hand.sum > 15 || (hand.hasAce && hand.sum >= 8 && hand.sum <= 11))
	return hand.score() <= 16 || (hand.hasAce && hand.sum <= 8)
}

func OpponentHitStrategy(hand *Hand) bool {
	return hand.score() <= 16 || (hand.hasAce && hand.sum <= 7)
}

func IsHit(hand *Hand) bool {
	return hand.score() < Target
}

func main() {
	wins := 0
	draws := 0
	losses := 0

	for i := 0; i < Rounds; i++ {
		switch PlayBlackjack(PlayerHitStrategy, OpponentHitStrategy) {
		case WIN:
			wins++
		case DRAW:
			draws++
		case LOSS:
			losses++
		}
	}

	fmt.Printf("Wins: %d (%f%%)\nDraws: %d (%f%%)\nLosses: %d (%f%%)\n",
		wins, float64(wins)/Rounds*100,
		draws, float64(draws)/Rounds*100,
		losses, float64(losses)/Rounds*100)
}
