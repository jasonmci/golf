package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestLoadCourse( t *testing.T) {
	courseName := "test-course"
	filePath := fmt.Sprintf("courses/%s.json", courseName)

	// Sample course data
	courseData := Course{
		Name: "Test Course",
		Holes: map[string]Hole{
			"1": {
				Par: 4,
				Outcomes: map[string]Outcome{
					"4": {W: "4a", I: "4B", P: "4X"},
				},
			},
		},
	}
	// Convert the sample course data to JSON and write it to the expected file
	fileContent, err := json.Marshal(courseData)
	if err != nil {
		t.Fatalf("Failed to marshal course data: %v", err)
	}
	
	err = os.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		t.Fatalf("Failed to write course file: %v", err)
	}

	loadedCourse := loadCourse(courseName)

	if loadedCourse.Name != courseData.Name {
		t.Errorf("Expected course name %s, got %s", courseData.Name, loadedCourse.Name)
	}

	if len(loadedCourse.Holes) != len(courseData.Holes) {
		t.Errorf("Expected %d holes, got %d", len(courseData.Holes), len(loadedCourse.Holes))
	}

	// check specific hole data

	hole, exists := loadedCourse.Holes["1"]
	if !exists {
		t.Error("Expected hole 1 to exist")
	} else {
		if hole.Par != 4 {
			t.Errorf("Expected par 4, got %d", hole.Par)
		}
	}

	defer func() {
		err = os.Remove(filePath)
		if err != nil {
			t.Fatalf("Failed to remove test file: %v", err)
		}
	}()

}

func TestDescribeHoleOutcome(t *testing.T) {

	// test struct for all describeHoleOutcome tests
	tests := []struct {
		outcome int
		par int
		expectedDescription string
	}{
		{2, 3, "Birdie"},
		{3, 5, "Eagle"},
		{2, 5, "Double Eagle"},
		{4, 4, "Par"},
		{5, 4, "Bogey"},
		{6, 4, "Double Bogey"},
		{7, 4, "Triple Bogey"},
		{11, 2, "Quadruple Bogey or Worse"},
	}

	for _, tt := range tests {
		description := describeHoleOutcome(tt.outcome, tt.par)

		if description != tt.expectedDescription {
			t.Errorf("Expected description %s, got %s", tt.expectedDescription, description)
		}
	}
}


func TestHoleOutcome(t *testing.T) {

	// set course conditions
	conditions := &CourseConditions{
		Daily: []string{"X", "x", "a"},
	}

	loadedPlayer := PlayerCard{
		Name: "Test Jack Nicklaus",
		Rolls: map[string]RollOutcome{},
		QuickPlay: 9,
	}

	
	course := Course{
		Name: "Test Pebble Beach",
		Holes: map[string]Hole{
			"1": {
				Par: 4,
				Outcomes: map[string]Outcome{
					"12": {W: "4a", I: "5A", P: "6X"},
					"34": {W: "3B", I: "4d", P: "5"},
				},
			},
			"2": {
				Par: 5,
				Outcomes: map[string]Outcome{
					"56": {W: "5x", I: "6D", P: "7"},
				},
			},
		},
	}

	tests := []struct {
		holeNumber string
		playerOutcome string
		thirdDie int
		expectedInt int
		expectedDesc string
		shouldError bool
	}{
		{"1", "12", 1, 5, "Bogey", false},
		{"1", "12", 3, 5, "Bogey", false},
		{"1", "12", 5, 4, "Par", false},
		{"1", "34", 1, 3, "Birdie", false},
		{"1", "34", 3, 4, "Par", false},
		{"1", "34", 5, 5, "Bogey", false},
		{"2", "56", 1, 5, "Par", false},
		{"2", "56", 3, 6, "Bogey", false},
		{"2", "56", 5, 7, "Double Bogey", false},
		{"3", "56", 5, 0, "", true},
	}

	for _, tt := range tests {
		outcomeInt, description, err := holeOutcome(course, loadedPlayer, conditions, tt.holeNumber, tt.playerOutcome, tt.thirdDie)

		if (err != nil) != tt.shouldError {
			t.Errorf("holeOutcome(%s, %s, %d): expected error = %v, got %v",
				tt.holeNumber, tt.playerOutcome, tt.thirdDie, tt.shouldError, err != nil)
		}

		if outcomeInt != tt.expectedInt {
			t.Errorf("holeOutcome(%s, %s, %d): expected outcomeInt = %d, got %d",
				tt.holeNumber, tt.playerOutcome, tt.thirdDie, tt.expectedInt, outcomeInt)
		}

		if description != tt.expectedDesc {
			t.Errorf("holeOutcome(%s, %s, %d): expected description = %s, got %s",
				tt.holeNumber, tt.playerOutcome, tt.thirdDie, tt.expectedDesc, description)
		}


	}

}
