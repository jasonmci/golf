package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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