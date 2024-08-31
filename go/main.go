package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Load the course data
	loadedCourse := loadCourse("cypress-point")
	
	// create a slice of player cards
	var playerCards = []PlayerCard{}

	playerCards = append(playerCards, loadPlayerCard("phil-mickelson"))
	playerCards = append(playerCards, loadPlayerCard("arnold-palmer"))
	playerCards = append(playerCards, loadPlayerCard("tiger-woods"))
	playerCards = append(playerCards, loadPlayerCard("jack-nicklaus"))

	// determine the course conditions
	courseCondition, wind, rough := rollDice()
	dailyConditions := courseModifiers(courseCondition, wind, rough)

	fmt.Printf("Course Conditions: %+v\n", dailyConditions)

	// create a slice of scorecards for each player
	var scorecards []Scorecard
	for _, player := range playerCards {
		scorecards = append(scorecards, createScorecard(player.Name, "🇺🇸", loadedCourse))
	}

  		// play a round of 18 holes with one player
		for i := 1; i <= 18; i++ {
			for j, player := range playerCards {
				d1, d2, d3 := rollDice()

				// convert int i to a string
				stringInt := strconv.Itoa(i)
				playerOutcome, _ := playerOutcome(player, d1, d2, d3) // ignore the modifier for now
				holeOutcome, _, _ := holeOutcome(loadedCourse, player, dailyConditions, stringInt, playerOutcome, d3)
				updateScorecard(&scorecards[j], i, holeOutcome)
				printScorecard(scorecards[j])
			}
			LeaderBoard(scorecards)
		}	
}
