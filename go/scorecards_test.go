package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCreateScorecard(t *testing.T) {
	// Define a sample Course with 18 holes
	course := Course{
		Name: "Test Course",
		Holes: map[string]Hole{
			"1":  {Par: 4},
			"2":  {Par: 3},
			"3":  {Par: 5},
			"4":  {Par: 4},
			"5":  {Par: 4},
			"6":  {Par: 3},
			"7":  {Par: 4},
			"8":  {Par: 5},
			"9":  {Par: 4},
			"10": {Par: 4},
			"11": {Par: 3},
			"12": {Par: 5},
			"13": {Par: 4},
			"14": {Par: 4},
			"15": {Par: 3},
			"16": {Par: 5},
			"17": {Par: 4},
			"18": {Par: 4},
		},
	}

	// Create the scorecard
	scorecard := createScorecard("John Doe", "🇺🇸", course)

	// Validate the scorecard data
	if scorecard.GolferName != "John Doe" {
		t.Errorf("Expected golfer name to be 'John Doe', got %s", scorecard.GolferName)
	}

	if scorecard.CountryFlag != "🇺🇸" {
		t.Errorf("Expected country flag to be '🇺🇸', got %s", scorecard.CountryFlag)
	}

	expectedHoles := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	if len(scorecard.Holes) != 18 {
		t.Errorf("Expected 18 holes, got %d", len(scorecard.Holes))
	}

	for i, hole := range scorecard.Holes {
		if hole != expectedHoles[i] {
			t.Errorf("Expected hole %d, got %d", expectedHoles[i], hole)
		}
	}

	expectedPar := []int{4, 3, 5, 4, 4, 3, 4, 5, 4, 4, 3, 5, 4, 4, 3, 5, 4, 4}
	if len(scorecard.Par) != 18 {
		t.Errorf("Expected 18 par values, got %d", len(scorecard.Par))
	}

	for i, par := range scorecard.Par {
		if par != expectedPar[i] {
			t.Errorf("Expected par %d for hole %d, got %d", expectedPar[i], i+1, par)
		}
	}

	if len(scorecard.Strokes) != 0 {
		t.Errorf("Expected strokes to be initialized to an empty slice, got %v", scorecard.Strokes)
	}

	if len(scorecard.RunningTotal) != 0 {
		t.Errorf("Expected running total to be initialized to an empty slice, got %v", scorecard.RunningTotal)
	}

	if len(scorecard.RunningPar) != 0 {
		t.Errorf("Expected running par to be initialized to an empty slice, got %v", scorecard.RunningPar)
	}
}

func TestUpdateScorecard(t *testing.T) {
	// Define a sample Course with 18 holes
	course := Course{
		Name: "Test Course",
		Holes: map[string]Hole{
			"1": {Par: 4},
			"2": {Par: 3},
			"3": {Par: 5},
		},
	}

	// Create the initial scorecard
	scorecard := createScorecard("John Doe", "🇺🇸", course)

	// Test updating the scorecard for the first hole
	updateScorecard(&scorecard, 1, 4)
	if len(scorecard.Strokes) != 1 || scorecard.Strokes[0] != 4 {
		t.Errorf("Expected first stroke to be 4, got %v", scorecard.Strokes)
	}
	if len(scorecard.RunningTotal) != 1 || scorecard.RunningTotal[0] != 4 {
		t.Errorf("Expected running total to be 4, got %v", scorecard.RunningTotal)
	}
	if len(scorecard.RunningPar) != 1 || scorecard.RunningPar[0] != 4 {
		t.Errorf("Expected running par to be 4, got %v", scorecard.RunningPar)
	}

	// Test updating the scorecard for the second hole
	updateScorecard(&scorecard, 2, 3)
	if len(scorecard.Strokes) != 2 || scorecard.Strokes[1] != 3 {
		t.Errorf("Expected second stroke to be 3, got %v", scorecard.Strokes)
	}
	if len(scorecard.RunningTotal) != 2 || scorecard.RunningTotal[1] != 7 {
		t.Errorf("Expected running total to be 7, got %v", scorecard.RunningTotal)
	}
	if len(scorecard.RunningPar) != 2 || scorecard.RunningPar[1] != 7 {
		t.Errorf("Expected running par to be 7, got %v", scorecard.RunningPar)
	}

	// Test updating the scorecard for the third hole
	updateScorecard(&scorecard, 3, 5)
	if len(scorecard.Strokes) != 3 || scorecard.Strokes[2] != 5 {
		t.Errorf("Expected third stroke to be 5, got %v", scorecard.Strokes)
	}
	if len(scorecard.RunningTotal) != 3 || scorecard.RunningTotal[2] != 12 {
		t.Errorf("Expected running total to be 12, got %v", scorecard.RunningTotal)
	}
	if len(scorecard.RunningPar) != 3 || scorecard.RunningPar[2] != 12 {
		t.Errorf("Expected running par to be 12, got %v", scorecard.RunningPar)
	}

	// Test invalid hole number (should not update anything)
	updateScorecard(&scorecard, 19, 4)
	if len(scorecard.Strokes) != 3 {
		t.Errorf("Invalid hole update should not change strokes, expected 3, got %d", len(scorecard.Strokes))
	}
	if len(scorecard.RunningTotal) != 3 {
		t.Errorf("Invalid hole update should not change running total, expected 3, got %d", len(scorecard.RunningTotal))
	}
	if len(scorecard.RunningPar) != 3 {
		t.Errorf("Invalid hole update should not change running par, expected 3, got %d", len(scorecard.RunningPar))
	}
}

func TestPrintScorecard(t *testing.T) {
	// Define a sample Scorecard
	scorecard := Scorecard{
		GolferName:   "John Doe",
		CountryFlag:  "🇺🇸",
		Holes:        []int{1, 2, 3},
		Par:          []int{4, 4, 5},
		Strokes:      []int{3, 5, 5},
		RunningTotal: []int{3, 8, 13},
		RunningPar:   []int{4, 8, 13},
	}

	// Capture the output of the printScorecard function
	//var output bytes.Buffer
	fmt.Printf("%s %s\n", scorecard.GolferName, scorecard.CountryFlag)
	fmt.Println(strings.Repeat("-", 87))

	// Print holes
	fmt.Printf("%-10s", "Hole:")
	for _, hole := range scorecard.Holes {
		fmt.Printf("%3d ", hole)
	}
	fmt.Println()
	fmt.Println(strings.Repeat("-", 87))

	// Print par
	fmt.Printf("%-10s", "Par:")
	for _, p := range scorecard.Par {
		fmt.Printf("%3d ", p)
	}
	fmt.Println()

	// Print strokes
	fmt.Printf("%-10s", "Strokes:")
	for i, stroke := range scorecard.Strokes {
		if stroke < scorecard.Par[i] {
			fmt.Printf("\033[32m%3d\033[0m ", stroke)
		} else if stroke > scorecard.Par[i] {
			fmt.Printf("\033[31m%3d\033[0m ", stroke)
		} else {
			fmt.Printf("%3d ", stroke)
		}
	}
	fmt.Println()
	fmt.Println(strings.Repeat("-", 87))

	// Print running total
	fmt.Printf("%-10s", "Total:")
	for _, total := range scorecard.RunningTotal {
		fmt.Printf("%3d ", total)
	}
	fmt.Println()

	// Determine +, E, or - score
	fmt.Printf("%-10s", "Score:")

	// Get the running total from the Par slice
	for i, total := range scorecard.RunningTotal {
		if total-scorecard.RunningPar[i] == 0 {
			fmt.Printf("%3s ", "E")
		} else if total-scorecard.RunningPar[i] > 0 && total-scorecard.RunningPar[i] < 10 {
			fmt.Printf(" +%d ", total-scorecard.RunningPar[i])
		} else if total-scorecard.RunningPar[i] > 10 {
			fmt.Printf("+%2d ", total-scorecard.RunningPar[i])
		} else {
			fmt.Printf("%3d ", total-scorecard.RunningPar[i])
		}
	}
	fmt.Println()
}