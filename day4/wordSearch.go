package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFile = "./input.txt"

func main() {
	b, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("reading rawInput file %s: %v", inputFile, err)
	}
	rawInput := strings.Split(string(b), "\n")
	inputSlice := make([][]string, 0, len(rawInput))

	// Build the 2D slice
	for _, row := range rawInput {
		rowSlice := strings.Split(row, "")
		inputSlice = append(inputSlice, rowSlice)
	}

	matches := calculateWordCount(inputSlice)
	fmt.Printf("Total matches: %d\n", matches)
}

// indexPoint represents an index location in a 2D slice.
type indexPoint struct {
	rowIndex    int
	columnIndex int
}

// calculateWordCount iterates through the input 2D slice and calculates how many times the string "XMAS" appears.
func calculateWordCount(input [][]string) int {
	totalMatches := 0

	for rowIndex, row := range input {
		for columnIndex := range row {

			matches := checkIndexPointDirections(indexPoint{rowIndex: rowIndex, columnIndex: columnIndex},
				input, len(input), len(row))

			fmt.Printf("Row index: %d, Column index: %d, Matches: %d\n", rowIndex, columnIndex, matches)

			totalMatches += matches
		}
	}

	return totalMatches
}

// checkIndexPointDirections takes an indexPoint and iterates through all index directions looking for an "XMAS" string match.
func checkIndexPointDirections(index indexPoint, inputSlice [][]string, numRows, numColumns int) int {
	matches := 0

	// The 2D index directions
	indexDiffs := map[string][][]int{
		"up":         {{-1, 0}, {-2, 0}, {-3, 0}},
		"up-right":   {{-1, 1}, {-2, 2}, {-3, 3}},
		"right":      {{0, 1}, {0, 2}, {0, 3}},
		"down-right": {{1, 1}, {2, 2}, {3, 3}},
		"down":       {{1, 0}, {2, 0}, {3, -0}},
		"down-left":  {{1, -1}, {2, -2}, {3, -3}},
		"left":       {{0, -1}, {0, -2}, {0, -3}},
		"up-left":    {{-1, -1}, {-2, -2}, {-3, -3}},
	}

	// Check for consecutive string match: "XMAS"
	for _, diff := range indexDiffs {
		isValidPoint := isValidPointDirection(indexPoint{rowIndex: index.rowIndex, columnIndex: index.columnIndex}, numRows, numColumns)
		if isValidPoint && inputSlice[index.rowIndex][index.columnIndex] == "X" {

			isValidPoint = isValidPointDirection(indexPoint{rowIndex: index.rowIndex + diff[0][0], columnIndex: index.columnIndex + diff[0][1]}, numRows, numColumns)
			if isValidPoint && inputSlice[index.rowIndex+diff[0][0]][index.columnIndex+diff[0][1]] == "M" {

				isValidPoint = isValidPointDirection(indexPoint{rowIndex: index.rowIndex + diff[1][0], columnIndex: index.columnIndex + diff[1][1]}, numRows, numColumns)
				if isValidPoint && inputSlice[index.rowIndex+diff[1][0]][index.columnIndex+diff[1][1]] == "A" {

					isValidPoint = isValidPointDirection(indexPoint{rowIndex: index.rowIndex + diff[2][0], columnIndex: index.columnIndex + diff[2][1]}, numRows, numColumns)
					if isValidPoint && inputSlice[index.rowIndex+diff[2][0]][index.columnIndex+diff[2][1]] == "S" {
						matches++
					}
				}
			}
		}
	}

	return matches
}

// isValidPointDirection checks if an indexPoint is out of bounds based on the row and column size.
func isValidPointDirection(index indexPoint, numRows, numColumns int) bool {
	if index.rowIndex < 0 || index.rowIndex >= numRows {
		return false
	}
	if index.columnIndex < 0 || index.columnIndex >= numColumns {
		return false
	}
	return true
}
