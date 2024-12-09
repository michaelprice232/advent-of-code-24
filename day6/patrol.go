package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFile = "./input.txt"

type direction string

const (
	north direction = "north"
	east  direction = "east"
	south direction = "south"
	west  direction = "west"
)

// location holds information about the current location on the map and also the direction we are facing.
type location struct {
	rowIndex    int
	columnIndex int
	direction   direction
}

func main() {
	patrolMap, err := parseMap(inputFile)
	if err != nil {
		log.Fatalf("parsing input text file: %v", err)
	}

	loc, err := findStartLocation(patrolMap)
	if err != nil {
		log.Fatalf("finding start location: %v", err)
	}

	fmt.Printf("Start location: %#v\n", loc)

	result := mapRoute(patrolMap, loc)
	fmt.Printf("Number of steps taken: %d\n", result)
}

// parseMap parses the input file and builds a 2D patrol map after validating the contents.
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

// validPatrolMap validates a patrol map.
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
			return false, fmt.Sprintf("mismatching row lengths. expected %d but got %d on row %d", firstRowLength, len(row), rowIdx)
		}
	}

	return true, ""
}

// findStartLocation returns a location based on starting point on the map.
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

// mapRoute iterates over the 2D patrol map following the directions set out in the requirements until we come off the map.
// Returns the number of unique places which have been visited.
// See readme for the requirements around moving.
func mapRoute(patrolMap [][]string, startLocation location) int {
	// Keep track of unique locations which have been visited. Include the starting location
	visitedLocations := make([][]bool, len(patrolMap))
	for i := range visitedLocations {
		visitedLocations[i] = make([]bool, len(patrolMap[0]))
	}
	visitedLocations[startLocation.rowIndex][startLocation.columnIndex] = true
	distinctLocationsVisited := 1

	// Iterate until we have come off the map
	loc := startLocation
	for {
		finished, newLocation := nextLocation(patrolMap, loc)

		if finished {
			fmt.Printf("We are now off the map!\n")
			return distinctLocationsVisited
		}

		if !visitedLocations[newLocation.rowIndex][newLocation.columnIndex] {
			distinctLocationsVisited++
			visitedLocations[newLocation.rowIndex][newLocation.columnIndex] = true
		}

		fmt.Printf("New location: %#v\n", newLocation)
		loc = newLocation
	}
}

// nextLocation returns the next location that should be visited based on a passed location.
func nextLocation(patrolMap [][]string, currentLocation location) (bool, location) {
	proposedNewLocation := currentLocation

	switch proposedNewLocation.direction {
	case north:
		proposedNewLocation.rowIndex--
	case east:
		proposedNewLocation.columnIndex++
	case south:
		proposedNewLocation.rowIndex++
	case west:
		proposedNewLocation.columnIndex--
	}

	// Use bool to indicate that we have gone off map and so are finished.
	if !stillOnMap(len(patrolMap), len(patrolMap[0]), proposedNewLocation) {
		return true, location{}
	}

	// Check if we would end up hitting an obstacle, rotate if so
	if patrolMap[proposedNewLocation.rowIndex][proposedNewLocation.columnIndex] == "#" {
		fmt.Printf("We would hit an obsacle at row %d column %d, turning 90 degrees right instad\n", proposedNewLocation.rowIndex, proposedNewLocation.columnIndex)
		currentLocation.direction = rotateDirection(currentLocation.direction)
		return false, currentLocation
	}

	// No obstacle or end of map
	return false, proposedNewLocation
}

// stillOnMap returns true if the passed location is still within the bounds of the 2D patrol map.
func stillOnMap(rowLength, columnLength int, location location) bool {
	if location.rowIndex < 0 || location.rowIndex >= rowLength {
		return false
	}

	if location.columnIndex < 0 || location.columnIndex >= columnLength {
		return false
	}

	return true
}

// rotateDirection rotates the direction by 90 degrees to the right.
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
