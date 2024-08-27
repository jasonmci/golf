package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Hole represents the outcomes for all rolls on a single hole
type Hole struct {
	Par      int                `json:"par"`
	Outcomes map[string]Outcome `json:"outcomes"`
}

// Course represents a golf course with a name and a map of holes
type Course struct {
	Name  string          `json:"name"`
	Holes map[string]Hole `json:"holes"`
}

func loadCourse(name string) Course {
	filePath := fmt.Sprintf("courses/%s.json", name)
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var course Course
	err = json.Unmarshal(file, &course)
	if err != nil {
		log.Fatal(err)
	}

	return course
}

func holeOutcome(course Course, holeNumber string, playerOutcome string, thirdDie int) (int, string, error) {
	holeData, exists := course.Holes[holeNumber]
	if !exists {
		return 0, "", fmt.Errorf("hole %s does not exist", holeNumber)
	}
	
	var holeOutcome string
	switch thirdDie {
	case 1, 2:
		holeOutcome = holeData.Outcomes[playerOutcome].W
	case 3, 4:
		holeOutcome = holeData.Outcomes[playerOutcome].I
	case 5, 6:
		holeOutcome = holeData.Outcomes[playerOutcome].P
	}

	outcomeInt, err := strconv.Atoi(holeOutcome)
	if err != nil {
		fmt.Print("Error converting string to int")
		//return 0,"", "", fmt.Errorf("error converting string to int: %v", err)
	}

	description := describeHoleOutcome(outcomeInt, holeData.Par)
	return outcomeInt, description, nil
}

func describeHoleOutcome(outcomeInt, par int) string {
	switch outcomeInt - par {
	case 0:
		return "Par"
	case -1:
		return "Birdie"
	case -2:
		return "Eagle"
	case -3:
		return "Double Eagle"
	case 1:
		return "Bogey"
	case 2:
		return "Double Bogey"
	case 3:
		return "Triple Bogey"
	default:
		return "Quadruple Bogey or Worse"
	}
}