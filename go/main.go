package main

import (
	"fmt"
	"sort"
	"strconv"
)

type LeaderBoardEntry struct {
	Name       		string
	Country    		string
	LastHole   		int
	RunningTotal 	int
	Diff			int
}

func LeaderBoard(scorecards []Scorecard) {
	var leaderBoard []LeaderBoardEntry
	for _, scorecard := range scorecards {
		leaderBoard = append(leaderBoard, LeaderBoardEntry{
			Name:         scorecard.GolferName,
			Country:      scorecard.CountryFlag,
			LastHole:     len(scorecard.RunningTotal),
			RunningTotal: scorecard.RunningTotal[len(scorecard.RunningTotal)-1],
			Diff: 	      scorecard.RunningTotal[len(scorecard.RunningTotal)-1] - scorecard.RunningPar[len(scorecard.RunningPar)-1],
		})
	}

	sort.Slice(leaderBoard, func(i,j int) bool {
		return leaderBoard[i].RunningTotal < leaderBoard[j].RunningTotal
	})

	// Print the leaderboard
	fmt.Println("Leaderboard:")
	fmt.Printf("%-20s %-10s %-15s %-10s %10s\n", "Name", "Country", "Last Hole", "Score", "Diff")
	for _, entry := range leaderBoard {
		fmt.Printf("%-20s %-10s %-15d %-10d %10d\n", entry.Name, entry.Country, entry.LastHole, entry.RunningTotal, entry.Diff)
	}
}

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
