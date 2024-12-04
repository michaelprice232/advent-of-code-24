package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*
WIP - not working 100% yet!

1. Read each line into slice of string slices. Check that all lines are equal length.
2. Have a function to check which directions are accessible. Return slice of enums
3. Have another function to return a slice of 3 index points to check - use an offset + / - ?
4. Iterate over each index and call the two functions. Check for spelling and increment counter
*/

const inputFile = "./input.txt"

func main() {
	b, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("reading rawInput file %s: %v", inputFile, err)
	}
	rawInput := strings.Split(string(b), "\n")
	inputSlice := make([][]string, 0, len(rawInput))

	for _, row := range rawInput {
		rowSlice := strings.Split(row, "")
		inputSlice = append(inputSlice, rowSlice)
	}

	matches := calculateWordCount(inputSlice)
	fmt.Printf("Total matches: %d\n", matches)
}

func calculateWordCount(input [][]string) int {
	totalMatches := 0

	for rowIndex, row := range input {
		for columnIndex, character := range row {

			matches := checkIndexPointDirections(indexPoint{rowIndex: rowIndex, columnIndex: columnIndex},
				input, len(input), len(row))

			fmt.Printf("Row index: %d, Column index: %d, Character: %s, Matches: %d\n", rowIndex, columnIndex, character, matches)

			totalMatches += matches
		}
	}

	return totalMatches
}

type indexPoint struct {
	rowIndex    int
	columnIndex int
}

func isValidPointDirection(index indexPoint, numRows, numColumns int) bool {
	if index.rowIndex < 0 || index.rowIndex >= numRows {
		return false
	}
	if index.columnIndex < 0 || index.columnIndex >= numColumns {
		return false
	}
	return true
}

func checkIndexPointDirections(index indexPoint, inputSlice [][]string, numRows, numColumns int) int {
	matches := 0

	indexDiffs := map[string][][]int{
		"up":         {{-1, 0}, {-2, 0}, {-3, 0}},
		"up-right":   {{-1, 1}, {-2, 2}, {-3, 3}},
		"right":      {{0, 1}, {0, 2}, {0, 3}},
		"down-right": {{1, 1}, {2, 2}, {3, 3}},
		"down":       {{1, 0}, {2, 0}, {3, -0}},
		"down-left":  {{1, -1}, {2, -2}, {3, -3}},
		"left":       {{0, -1}, {0, -2}, {0, -3}},
		"up-left":    {{1, -1}, {2, -2}, {3, -3}},
	}

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

	// Check diagonal up left
	return matches
}
