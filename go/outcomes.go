package main

import (
	"math/rand"
	"time"
)

// RollOutcome represents the outcomes for W, I, and P for a particular roll
type RollOutcome struct {
	W string `json:"w"`
	I string `json:"i"`
	P string `json:"p"`
}

// Outcome represents the W, I, P outcomes for a particular shot
// this should be the one based on the course but we should check on that
type Outcome struct {
	W string `json:"w"`
	I string `json:"i"`
	P string `json:"p"`
}

// rollDice rolls three six-sided dice and returns their values.
func rollDice() (int, int, int) {
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
	die1 := rand.Intn(6) + 1
	die2 := rand.Intn(6) + 1
	die3 := rand.Intn(6) + 1
	return die1, die2, die3
}