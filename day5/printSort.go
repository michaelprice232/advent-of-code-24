package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
2 lots of parsing:
1. 2 string slices for the orders and updates list
2. pass them to separate function to parse into the proper data structures


For the page rule ordering number create a map which contains a slice of the numbers it must come before
Iterate through each update number in the update input and for the numbers BEFORE it check they aren't in the list for that number (== violation)

map[int][]int
1: 2,3,4

If we have iterated through all update numbers and there are no violations then it's valid

If it's valid calculate the middle value and append the value to the running total

*/

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

	for k, v := range orderings {
		fmt.Printf("%d: %d\n", k, v)
	}
	fmt.Println()
	for i, v := range updates {
		fmt.Printf("%d: %d\n", i, v)
	}
}

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
			return orderings, fmt.Errorf("converting string '%s' to int: %w", ordering[0], err)
		}

		// Check if the key already exists in the map. Initialise an empty slice value if not
		_, found := orderings[number]
		if found {
			orderings[number] = append(orderings[number], dependentNumber)
		} else {
			orderings[number] = []int{dependentNumber}
		}
	}

	return orderings, nil
}

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

func checkUpdate(update []int, orderings map[int][]int) bool {
	isValid := true

	for idx, number := range update {
		fmt.Printf("Checking index %d: %d\n", idx, number)

		//// We are already at the start of the slice. Nothing before it to process
		//if idx == 0 {
		//	break
		//}

		// Process all the elements which appear before this number and check if they should appear afterward
		for i := idx - 1; i > 0; i-- {
			currentNumber := update[i]
			// todo: check if the dependent number exists in the map
		}

	}

	return isValid
}
