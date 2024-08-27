package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {

	// Load the course data
	loadedCourse := loadCourse("cypress-point")

	// Load the player card
	loadedPlayer := loadPlayerCard("phil-mickelson")
	fmt.Printf("Loaded Player Card: %+v\n", loadedPlayer.Name)

	d1, d2, d3 := rollDice()
	roll := fmt.Sprintf("%d%d", d1, d2)

	// get the roll outcome for the player
	rollOutcome := loadedPlayer.Rolls[roll]

	// take the third die, if its a 1 or 2, the outcome is W, for 3 or 4, the outcome is I, for 5 or 6, the outcome is P
	thirdDie := d3
	var playerOutcome string
	if thirdDie == 1 || thirdDie == 2 {
		playerOutcome = rollOutcome.W
	}
	if thirdDie == 3 || thirdDie == 4 {
		playerOutcome = rollOutcome.I
	}
	if thirdDie == 5 || thirdDie == 6 {
		playerOutcome = rollOutcome.P
	}

	// if the player outcome has a letter in it, we need to strip it out and save it to the playerOutcomeModifier variable
	// and then replace the letter with a space
	var playerOutcomeModifier string
	if strings.Contains(playerOutcome, "p") {
		playerOutcomeModifier = "p"
		playerOutcome = strings.ReplaceAll(playerOutcome, "p", "")
	}
	if strings.Contains(playerOutcome, "t") {
		playerOutcomeModifier = "t"
		playerOutcome = strings.ReplaceAll(playerOutcome, "t", "")
	}
	if strings.Contains(playerOutcome, "r") {
		playerOutcomeModifier = "r"
		playerOutcome = strings.ReplaceAll(playerOutcome, "r", "")
	}
	if strings.Contains(playerOutcome, "m") {
		playerOutcomeModifier = "m"
		playerOutcome = strings.ReplaceAll(playerOutcome, "m", "")
	}
	if strings.Contains(playerOutcome, "s") {
		playerOutcomeModifier = "s"
		playerOutcome = strings.ReplaceAll(playerOutcome, "s", "")
	}
	if strings.Contains(playerOutcome, "d") {
		playerOutcomeModifier = "d"
		playerOutcome = strings.ReplaceAll(playerOutcome, "d", "")
	} else {
		playerOutcomeModifier = ""
	}

	fmt.Printf("Player Outcome: %s\n", playerOutcome)
	fmt.Printf("Player Outcome Modifier: %s\n", playerOutcomeModifier)

	//get the hole data for the first hole
	holeData := loadedCourse.Holes["1"]

	var holeOutcome string
	if thirdDie == 1 || thirdDie == 2 {
		holeOutcome = holeData.Outcomes[playerOutcome].W
	}
	if thirdDie == 3 || thirdDie == 4 {
		holeOutcome = holeData.Outcomes[playerOutcome].I
	}
	if thirdDie == 5 || thirdDie == 6 {
		holeOutcome = holeData.Outcomes[playerOutcome].P
	}
	fmt.Printf("Hole Outcome: %s\n", holeOutcome)

	// convert holeOutcome to a number

    outcomeInt, err := strconv.Atoi(holeOutcome)
    if err != nil {
        fmt.Println("Error converting string to int:", err)
        return
    }
	// describe the hole outcome
	if outcomeInt - holeData.Par == 0 {
		fmt.Println("You made par!")
	}
	if outcomeInt - holeData.Par == -1 {
		fmt.Println("You made birdie!")
	}
	if outcomeInt - holeData.Par == -2 {
		fmt.Println("You made eagle!")
	}
	if outcomeInt - holeData.Par == -3 {
		fmt.Println("You made double eagle!")
	}
	if outcomeInt - holeData.Par == 1 {
		fmt.Println("You made bogey!")
	}
	if outcomeInt - holeData.Par == 2 {
		fmt.Println("You made double bogey!")
	}
	if outcomeInt - holeData.Par == 3 {
		fmt.Println("You made triple bogey!")
	}
	if outcomeInt - holeData.Par == 4 {
		fmt.Println("You made quadruple bogey!")
	}
	if outcomeInt - holeData.Par >= 5 {
		fmt.Println("You made quintuple bogey or worse!")
	}

    scorecard := createScorecard(loadedPlayer.Name, "🇺🇸", loadedCourse)

    // Update the scorecard as the golfer completes each hole
    updateScorecard(&scorecard, 1, outcomeInt)
    updateScorecard(&scorecard, 2, 5)
    updateScorecard(&scorecard, 3, 3)
    updateScorecard(&scorecard, 4, 23)
	updateScorecard(&scorecard, 5, 4)
	updateScorecard(&scorecard, 6, 2)

    // Print the updated scorecard
    fmt.Println("\nUpdated Scorecard:")
    printScorecard(scorecard)
}
