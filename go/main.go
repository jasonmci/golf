package main

import (
	"fmt"
)

func main() {

	// Load the course data
	loadedCourse := loadCourse("cypress-point")
	
	// Load the player card
	loadedPlayer := loadPlayerCard("phil-mickelson")
	fmt.Printf("Loaded Player Card: %+v\n", loadedPlayer.Name)
	
	d1, d2, d3 := rollDice()
	
	playerOutcome, _ := playerOutcome(loadedPlayer, d1, d2, d3) // ignore the modifier for now
	
	holeOutcome, description, err := holeOutcome(loadedCourse, "1", playerOutcome, d3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Hole Outcome: %d\n", holeOutcome)
    fmt.Println(description)

    scorecard := createScorecard(loadedPlayer.Name, "🇺🇸", loadedCourse)

    // Update the scorecard as the golfer completes each hole
    updateScorecard(&scorecard, 1, holeOutcome)
    updateScorecard(&scorecard, 2, 5)
    updateScorecard(&scorecard, 3, 3)
    updateScorecard(&scorecard, 4, 3)
	updateScorecard(&scorecard, 5, 4)
	updateScorecard(&scorecard, 6, 4)

    // Print the updated scorecard
    fmt.Println("\nUpdated Scorecard:")
    printScorecard(scorecard)
}
