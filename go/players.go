package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// PlayerCard represents the player card with a name and rolls
type PlayerCard struct {
	Name      string                 `json:"name"`
	CountryFlag string                 `json:"countryFlag"`
	QuickPlay int                    `json:"qp"`
	Rolls     map[string]RollOutcome `json:"rolls"`
}

func playerOutcome(player PlayerCard, d1, d2, d3 int) (string, string) {
	roll := fmt.Sprintf("%d%d", d1, d2)
	rollOutcome := player.Rolls[roll]

	var outcome string
	switch d3 {
	case 1, 2:
		outcome = rollOutcome.W
	case 3, 4:
		outcome = rollOutcome.I
	case 5, 6:
		outcome = rollOutcome.P
	}

	var modifier string
	for _, letter := range []string{"p", "t", "r", "m", "s", "d"} {
		if strings.Contains(outcome, letter) {
			modifier = letter
			outcome = strings.ReplaceAll(outcome, letter, "")
			break
		}
	}
	 return outcome, modifier
}

func loadPlayerCard(name string) PlayerCard {
	filePath := fmt.Sprintf("player-cards/%s.json", name)
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var player PlayerCard
	err = json.Unmarshal(file, &player)
	if err != nil {
		log.Fatal(err)
	}

	return player
}