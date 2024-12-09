package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*
- Parse as 2D string slice
- Function which calculates the next position
  - current position + direction. Store as variables
  - next position based on direction
    - check if free or rotate
- increment counter for each movement block
- check if off the map == finish
*/

const inputFile = "./input.txt"

func main() {
	patrolMap, err := parseMap(inputFile)
	if err != nil {
		log.Fatalf("parsing input text file: %v", err)
	}

	//for rowIdx, rowLocation := range patrolMap {
	//	for columnIdx, character := range rowLocation {
	//		fmt.Printf("Row: %d, Column: %d, Character: %s\n", rowIdx, columnIdx, character)
	//	}
	//}

	loc, err := findStartLocation(patrolMap)
	if err != nil {
		log.Fatalf("finding start location: %v", err)
	}

	fmt.Printf("Start location: %#v\n", loc)

	result := mapRoute(patrolMap, loc)
	fmt.Printf("Number of steps taken: %d\n", result)
}

func mapRoute(patrolMap [][]string, startLocation location) int {
	// Determine what direction to move in and calculate new index
	// If new index is off the map - finished
	// If new index is empty (.) - move into it
	// If new index is an obstacle (#) - turn 90 degrees
	// Perform on a loop until off the map

	movesCount := 0
	for {
		finished, newLocation := nextLocation(patrolMap, startLocation)
		if finished {
			fmt.Printf("We are now off the map!")
			return movesCount
		}

		movesCount++
		fmt.Printf("New location: %#v\n", newLocation)

	}
}

func nextLocation(patrolMap [][]string, currentLocation location) (bool, location) {
	proposedNewLocation := currentLocation

	switch proposedNewLocation.direction {
	case north:
		proposedNewLocation.columnIndex--
	case east:
		proposedNewLocation.rowIndex++
	case south:
		proposedNewLocation.columnIndex++
	case west:
		proposedNewLocation.rowIndex--
	}

	// Use bool to indicate that we have gone off map and so are finished
	if !stillOnMap(len(patrolMap[0]), len(patrolMap), proposedNewLocation) {
		return true, location{}
	}

	// Check if we would end up hitting an obstacle, rotate if so
	if patrolMap[proposedNewLocation.rowIndex][proposedNewLocation.columnIndex] == "#" {
		fmt.Printf("We would hit an obsacle at row %d column %d, turning instad\n", proposedNewLocation.rowIndex, proposedNewLocation.columnIndex)
		currentLocation.direction = rotateDirection(currentLocation.direction)
		return false, currentLocation
	}

	return false, proposedNewLocation
}

func rotateDirection(direction direction) direction {
	switch direction {
	case north:
		return east
	case east:
		return south
	case south:
		return west
	case west:
		return north
	}

	return direction
}

func stillOnMap(rowLength, columnLength int, location location) bool {
	if location.rowIndex < 0 || location.rowIndex >= rowLength {
		return false
	}

	if location.columnIndex < 0 || location.columnIndex >= columnLength {
		return false
	}

	return true
}

type direction string

const (
	north direction = "north"
	east  direction = "east"
	south direction = "south"
	west  direction = "west"
)

type location struct {
	rowIndex    int
	columnIndex int
	direction   direction
}

func findStartLocation(patrolMap [][]string) (location, error) {
	foundLocations := 0
	l := location{}

	for rowIdx, row := range patrolMap {
		for columnIdx, character := range row {
			switch character {
			case "^":
				l.rowIndex = rowIdx
				l.columnIndex = columnIdx
				l.direction = north
				foundLocations++
			case ">":
				l.rowIndex = rowIdx
				l.columnIndex = columnIdx
				l.direction = east
				foundLocations++
			case "v":
				l.rowIndex = rowIdx
				l.columnIndex = columnIdx
				l.direction = south
				foundLocations++
			case "<":
				l.rowIndex = rowIdx
				l.columnIndex = columnIdx
				l.direction = west
				foundLocations++
			}
		}
	}

	if foundLocations != 1 {
		return l, fmt.Errorf("expected exactly 1 start location, but got %d", foundLocations)
	}

	return l, nil
}

func parseMap(inputFile string) ([][]string, error) {
	b, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("reading input text file: %v", err)
	}

	rawRows := strings.Split(string(b), "\n")

	patrolMap := make([][]string, 0, len(rawRows))

	for _, r := range rawRows {
		row := strings.Split(r, "")
		patrolMap = append(patrolMap, row)
	}

	valid, reason := validPatrolMap(patrolMap)
	if !valid {
		return nil, fmt.Errorf("validating patrol map failure: %s", reason)
	}

	return patrolMap, nil
}

func validPatrolMap(patrolMap [][]string) (bool, string) {
	firstRowLength := len(patrolMap[0])

	for rowIdx, row := range patrolMap {
		// Check for invalid characters
		for columnIdx, character := range row {
			if character != "." && character != "#" && character != "^" && character != ">" && character != "v" && character != "<" {
				return false, fmt.Sprintf("invalid chaacter '%s' at row %d column %d", character, rowIdx, columnIdx)
			}
		}

		// Check that all the row lengths are the same
		if len(row) != firstRowLength {
			return false, fmt.Sprintf("mismathing row lengths. expected %d but got %d on row %d", firstRowLength, len(row), rowIdx)
		}
	}

	return true, ""
}
