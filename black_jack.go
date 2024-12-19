package main

const (
	WIN = iota
	DRAW
	LOSS
)

type Hand struct {
	sum      int
	aceCount int
}

type HitStrategy func(Hand) bool

func play_blackjack(player_strategy HitStrategy, opponent_strategy HitStrategy) {

}
