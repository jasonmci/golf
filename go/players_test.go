package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestLoadPlayerCard(t *testing.T) {
	// Set up the actual file location as expected by the original function
	playerName := "test-gary-player"
	filePath := fmt.Sprintf("player-cards/%s.json", playerName)

	// Sample player data
	playerData := PlayerCard{
		Name: "Test Player",
		Rolls: map[string]RollOutcome{
			"11": {W: "12", I: "24", P: "1"},
		},
	}

	// Convert the sample player data to JSON and write it to the expected file
	fileContent, err := json.Marshal(playerData)
	if err != nil {
		t.Fatalf("Failed to marshal player data: %v", err)
	}

	err = os.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		t.Fatalf("Failed to write player card file: %v", err)
	}

	// Load the player card using the actual function
	loadedPlayer := loadPlayerCard(playerName)
	if err != nil {
		t.Fatalf("Failed to load player card: %v", err)
	}

	// Validate the loaded player data
	if loadedPlayer.Name != playerData.Name {
		t.Errorf("Expected player name %s, got %s", playerData.Name, loadedPlayer.Name)
	}

	if len(loadedPlayer.Rolls) != len(playerData.Rolls) {
		t.Errorf("Expected %d rolls, got %d", len(playerData.Rolls), len(loadedPlayer.Rolls))
	}

	// Clean up the test file
	err = os.Remove(filePath)
	if err != nil {
		t.Fatalf("Failed to remove test file: %v", err)
	}

}

func TestPlayerOutcome(t *testing.T) {
	player := PlayerCard{
		Name: "Test Jack Nicklaus",
		Rolls: map[string]RollOutcome{
			"11": {W: "12", I: "24t", P: "1"},
			"12": {W: "15m", I: "18", P: "20"},
			"13": {W: "5p", I: "10", P: "14"},
		},
	}

	// define the test cases
	tests := []struct {
		d1, d2, d3 int
		expectedOutcome string
		expectedModifier string
	}{
		{1, 1, 1, "12", ""},
		{1, 1, 4, "24", "t"},
		{1, 2, 2, "15", "m"},
		{1, 2, 5, "20", ""},
		{1, 3, 1, "5", "p"},
		{1, 3, 4, "10", ""},
	}

	for _, tt := range tests {
		outcome, modifier := playerOutcome(player, tt.d1, tt.d2, tt.d3)

		if outcome != tt.expectedOutcome {
			t.Errorf("Expected outcome %s, got %s", tt.expectedOutcome, outcome)
		}

		if modifier != tt.expectedModifier {
			t.Errorf("Expected modifier %s, got %s", tt.expectedModifier, modifier)
		}
	}
}