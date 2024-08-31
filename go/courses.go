package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

type CourseConditions struct {
	Daily []string
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


// This is a basic set of modifiers that can eventaully be set in the course data json
func courseModifiers(courseCondition int, wind int, rough int) *CourseConditions {

	dailyConditions := []string{"X", "x"}

	// write to daily conditions struct

	switch courseCondition {
	case 1:
		// append "A" to dailyConditions
		dailyConditions = append(dailyConditions, "A")
	case 3,4,5,6:
		dailyConditions = append(dailyConditions, "b")
	}

	switch wind {
	case 6:
		dailyConditions = append(dailyConditions, "c")
	}

	switch rough {
	case 1,2,3,4,5,6:
		// append nothing	
	}
	
	return &CourseConditions{
		Daily: dailyConditions,
	}
}

func holeOutcome(course Course, loadedPlayer PlayerCard, dailyConditions *CourseConditions, holeNumber string, playerOutcome string, thirdDie int) (int, string, error) {
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
	
	var modifier string
	for _, letter := range []string{"A", "a", "B", "B", "C", "c", "D", "d", "X", "x"} {
		if strings.Contains(holeOutcome, letter) {
			modifier = letter
			holeOutcome = strings.ReplaceAll(holeOutcome, letter, "")
			
		}
	}
	
	outcomeInt, err := strconv.Atoi(holeOutcome)
	if err != nil {
		fmt.Print("Error converting string to int")
	}
	
	for _, v := range dailyConditions.Daily {
		// check of modifier matches v
		if modifier == v {
			fmt.Printf("Matched: %s\n", v)
			if modifier == "A" || modifier == "B" || modifier == "C" || modifier == "D" {  
				outcomeInt -= 1 
			}
			if modifier == "a" || modifier == "b" || modifier == "c" || modifier == "d" {  
				outcomeInt += 1 
			}
			if modifier == "X" && outcomeInt >= loadedPlayer.QuickPlay {
				outcomeInt = holeData.Par - 1 
			} else if modifier == "X" && outcomeInt < loadedPlayer.QuickPlay {
				outcomeInt = holeData.Par
			}

			if modifier == "x" && outcomeInt >= loadedPlayer.QuickPlay {
				outcomeInt = holeData.Par + 1
			} else if modifier == "x" && outcomeInt < loadedPlayer.QuickPlay {
				outcomeInt = holeData.Par
			}

		}
	}

	description := describeHoleOutcome(outcomeInt, holeData.Par)
	return outcomeInt, description, nil
}

// provide the outcome int and the modifier to get the description
func describeHoleOutcome(outcomeInt int, par int) string {

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