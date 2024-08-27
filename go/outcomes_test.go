package main

import (
	"testing"
)

func TestRollDice(t *testing.T) {
	for i := 0; i < 100; i++ {
		die1, die2, die3 := rollDice()
		
		if die1 < 1 || die1 > 6 {
			t.Errorf("die1 out of range: %d", die1)
		}

		if die2 < 1 || die2 > 6 {
			t.Errorf("die2 out of range: %d", die2)
		}

		if die3 < 1 || die3 > 6 {
			t.Errorf("die3 out of range: %d", die3)
		}
		
	}
}