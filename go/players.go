package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// PlayerCard represents the player card with a name and rolls
type PlayerCard struct {
	Name      string                 `json:"name"`
	CountryFlag string                 `json:"countryFlag"`
	QuickPlay int                    `json:"qp"`
	Rolls     map[string]RollOutcome `json:"rolls"`
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