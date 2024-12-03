package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const inputFile = "./input.txt"

func main() {
	b, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("reading input file %s: %v", inputFile, err)
	}
	input := string(b)

	// Create a regular expression which matches 'mul(x,y)'. x and y can be 1->3 digits.
	// Use the inner brackets to create two capture groups, which extracts the digits separately.
	// See examples:
	// https://pkg.go.dev/regexp#Regexp.FindAllStringSubmatch
	// https://pkg.go.dev/regexp/syntax#section-sourcefiles (syntax)
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	count, err := calculateSum(input, re)
	if err != nil {
		log.Fatalf("calculating result: %v", err)
	}

	if count == 0 {
		fmt.Println("There were no results")
		os.Exit(0)
	}
	fmt.Printf("Result: %d\n", count)
}

// calculateSum calculates the sum of x*y in all the matching regular expressions, as described above.
func calculateSum(input string, exp *regexp.Regexp) (int, error) {
	total := 0
	results := exp.FindAllStringSubmatch(input, -1)

	if len(results) == 0 {
		return total, nil
	}

	for _, r := range results {
		if len(r) != 3 {
			return total, fmt.Errorf("expected 3 results per match. 1 for the full regular expression and 2 capture groups")
		}

		// Extract the two multiplication digits
		digitOne, err := strconv.Atoi(r[1])
		if err != nil {
			return total, fmt.Errorf("converting string %s to an int: %v", r[1], err)
		}
		digitTwo, err := strconv.Atoi(r[2])
		if err != nil {
			return total, fmt.Errorf("converting string %s to an int: %v", r[2], err)
		}

		total += digitOne * digitTwo
	}

	return total, nil
}
