package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const inputFile = "./input.txt"

func main() {
	b, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("reading input text file: %v", err)
	}

	raw := strings.Split(string(b), "\n")
	rawOrderings, rawUpdates := parseRaw(raw)

	orderings, err := parseOrderings(rawOrderings)
	if err != nil {
		log.Fatalf("parsing orderings input: %v", err)
	}

	updates, err := parseUpdates(rawUpdates)
	if err != nil {
		log.Fatalf("parsing updates input: %v", err)
	}

	result := processUpdates(updates, orderings)
	fmt.Printf("Total of the middle numbers for all valid updates: %d\n", result)
}

// parseRaw parses the text input into two slices for later processing.
// See readme and input.txt for sample input.
func parseRaw(input []string) ([]string, []string) {
	orderings := make([]string, 0)
	updates := make([]string, 0)

	for _, line := range input {
		if strings.Contains(line, "|") {
			orderings = append(orderings, line)
		}
		if strings.Contains(line, ",") {
			updates = append(updates, line)
		}
	}

	return orderings, updates
}

// parseOrderings parses the orderings slice and builds a map containing the numbers which it should appear before.
func parseOrderings(rawOrderings []string) (map[int][]int, error) {
	orderings := make(map[int][]int)

	for _, o := range rawOrderings {
		ordering := strings.Split(o, "|")
		if len(ordering) != 2 {
			return orderings, fmt.Errorf("expected 2 components for the ordering split for '%s', got %d", o, len(ordering))
		}

		number, err := strconv.Atoi(ordering[0])
		if err != nil {
			return orderings, fmt.Errorf("converting string '%s' to int: %w", ordering[0], err)
		}
		dependentNumber, err := strconv.Atoi(ordering[1])
		if err != nil {
			return orderings, fmt.Errorf("converting string '%s' to int: %w", ordering[1], err)
		}

		// Check if the key already exists in the map. Initialise slice if not
		_, found := orderings[number]
		if found {
			orderings[number] = append(orderings[number], dependentNumber)
		} else {
			orderings[number] = []int{dependentNumber}
		}
	}

	return orderings, nil
}

// parseUpdates parses the updates slice and converts them into int slices.
func parseUpdates(rawUpdates []string) ([][]int, error) {
	updates := make([][]int, 0, len(rawUpdates))

	for _, u := range rawUpdates {
		singleUpdateRaw := strings.Split(u, ",")
		singleUpdate := make([]int, 0, len(singleUpdateRaw))

		for _, updateStr := range singleUpdateRaw {
			number, err := strconv.Atoi(updateStr)
			if err != nil {
				return updates, fmt.Errorf("converting string '%s' to int: %w", updateStr, err)
			}
			singleUpdate = append(singleUpdate, number)
		}

		updates = append(updates, singleUpdate)
	}

	return updates, nil
}

// processUpdates processes each update and check that they are valid and therefore appear in the correct order.
func processUpdates(updates [][]int, orderings map[int][]int) int {
	result := 0
	for _, update := range updates {
		isValid, reason := checkUpdate(update, orderings)
		if isValid {
			middleNumber := len(update) / 2
			result += update[middleNumber]
		} else {
			fmt.Printf("Invalid update %v: %s\n", update, reason)
		}
	}

	return result
}

// checkUpdate checks an individual update to see if its valid.
func checkUpdate(update []int, orderings map[int][]int) (bool, string) {
	isValid := true
	reason := ""

	for idx, number := range update {
		for i := idx - 1; i >= 0; i-- {
			currentNumber := update[i]
			if slices.Contains(orderings[number], currentNumber) {
				reason = fmt.Sprintf("Number %d appears before %d when it shouldn't", currentNumber, number)
				isValid = false
			}
		}
	}

	return isValid, reason
}
