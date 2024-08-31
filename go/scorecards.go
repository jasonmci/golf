package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Scorecard struct {
	GolferName 		string
	CountryFlag 	string
	Holes 			[]int
	Par 			[]int
	Strokes 		[]int
	RunningTotal    []int
	RunningPar		[]int
}

func createScorecard(golferName string, countryFlag string, course Course) Scorecard {

	var strokes []int
	var runningTotal []int
	var runningPar []int

	var par []int

	for i := 1; i <= 18; i++ {
		num := strconv.Itoa(i)
		par = append(par, course.Holes[num].Par)
	}

	return Scorecard{
		GolferName: golferName,
		CountryFlag: countryFlag,
		Holes: []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18},
		Par: par,
		Strokes: strokes,
		RunningTotal: runningTotal,
		RunningPar: runningPar,
	}
}

func printScorecard(scorecard Scorecard) {
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

	// determine +, E, or - score
	fmt.Printf("%-10s", "Score:")

	// get the running total from the Par slice
	for i, total := range scorecard.RunningTotal {
		if total - scorecard.RunningPar[i] == 0 {
			fmt.Printf("%3s ", "E")
		} else if total - scorecard.RunningPar[i] > 0  && total - scorecard.RunningPar[i] < 10 {
			fmt.Printf(" +%d ", total - scorecard.RunningPar[i])
		} else if total - scorecard.RunningPar[i] > 10 {
			fmt.Printf("+%2d ", total - scorecard.RunningPar[i])
		} else {
			fmt.Printf("%3d ", total - scorecard.RunningPar[i])
		}
	}
	fmt.Println()
}

func updateScorecard(scorecard *Scorecard, hole int, strokes int){
	if hole < 1 || hole > len(scorecard.Holes) {
		fmt.Println("Invalid hole number")
		return
	}
	if len(scorecard.Strokes) == 0 {
		scorecard.RunningTotal = append(scorecard.RunningTotal, strokes)

		scorecard.RunningPar = append(scorecard.RunningPar, scorecard.Par[hole-1])
	} else {
		previousTotal := scorecard.RunningTotal[len(scorecard.RunningTotal)-1]
		previousTotal += strokes
		scorecard.RunningTotal = append(scorecard.RunningTotal, previousTotal)

		previousPar := scorecard.RunningPar[len(scorecard.RunningPar)-1]
		previousPar += scorecard.Par[hole-1]
		scorecard.RunningPar = append(scorecard.RunningPar, previousPar)
	}

	scorecard.Strokes = append(scorecard.Strokes, strokes)
}

type LeaderBoardEntry struct {
	Name       		string
	Country    		string
	LastHole   		int
	RunningTotal 	int
	Diff			string
}

func LeaderBoard(scorecards []Scorecard) {
	var leaderBoard []LeaderBoardEntry
	for _, scorecard := range scorecards {

		// set diff to have a + for all over par scores and E for even
		var mydiff string
		diffInt  := scorecard.RunningTotal[len(scorecard.RunningTotal)-1] - scorecard.RunningPar[len(scorecard.RunningPar)-1]
		if diffInt > 0 {
			mydiff = "+" + strconv.Itoa(diffInt)
		} else if diffInt < 0 {
			mydiff = strconv.Itoa(diffInt)
		} else if diffInt == 0 {
			mydiff = "E"
		}

		leaderBoard = append(leaderBoard, LeaderBoardEntry{
			Name:         scorecard.GolferName,
			Country:      scorecard.CountryFlag,
			LastHole:     len(scorecard.RunningTotal),
			RunningTotal: scorecard.RunningTotal[len(scorecard.RunningTotal)-1],
			Diff: 	      mydiff,
		})
	}

	sort.Slice(leaderBoard, func(i,j int) bool {
		return leaderBoard[i].RunningTotal < leaderBoard[j].RunningTotal
	})

	// Print the leaderboard
	fmt.Println("Leaderboard:")
	fmt.Printf("%-20s %-10s %-15s %-10s %10s\n", "Name", "Country", "Last Hole", "Score", "Diff")
	for _, entry := range leaderBoard {
		fmt.Printf("%-20s %-10s %-15d %-10d %10s\n", entry.Name, entry.Country, entry.LastHole, entry.RunningTotal, entry.Diff)
	}
}