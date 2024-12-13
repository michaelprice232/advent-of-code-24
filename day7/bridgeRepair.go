package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "./input.txt"

func main() {
	equations, err := parseConfig(inputFile)
	if err != nil {
		log.Fatalf("parsing config from %s: %v", inputFile, err)
	}

	fmt.Printf("Total = %d\n", calculateResults(equations))
}

// equation holds an individual equation.
type equation struct {
	target    int
	operands  []int
	operators []string
	valid     bool
}

// parseConfig return a validated slice of equation for later processing.
func parseConfig(inputFilePath string) ([]equation, error) {
	b, err := os.ReadFile(inputFilePath)
	if err != nil {
		return []equation{}, fmt.Errorf("reading file from %s: %v", inputFile, err)
	}
	if len(b) == 0 {
		return []equation{}, fmt.Errorf("expected at least 1 line of input")
	}

	rows := strings.Split(string(b), "\n")
	results := make([]equation, 0, len(rows))

	for _, row := range rows {
		e := equation{
			operands:  make([]int, 0),
			operators: make([]string, 0),
		}

		rowSplit := strings.Split(row, ":")
		if len(rowSplit) != 2 {
			return []equation{}, fmt.Errorf("expected only 2 components separated by : but got %d. line: %s", len(rowSplit), row)
		}

		target, err := strconv.Atoi(rowSplit[0])
		if err != nil {
			return []equation{}, fmt.Errorf("converting string to int: %w", err)
		}
		e.target = target

		operands := strings.Split(strings.TrimSpace(rowSplit[1]), " ")
		if len(operands) < 2 {
			return []equation{}, fmt.Errorf("expected at least 2 operands but got %d. line: %s", len(operands), row)
		}
		for _, operand := range operands {
			o, err := strconv.Atoi(operand)
			if err != nil {
				return []equation{}, fmt.Errorf("converting string to int: %w", err)
			}
			e.operands = append(e.operands, o)
		}

		// Populate the operators with all possible combinations based on the operands
		e.operators = generateVariations(len(e.operands) - 1)
		if len(e.operators) == 0 {
			return []equation{}, fmt.Errorf("expected at least 1 operator but got 0. line: %s", row)
		}

		results = append(results, e)
	}

	return results, nil
}

// Generate all the possible combinations of operators (* or +) based on an arbitrary number of operands.
// https://en.wikipedia.org/wiki/Cartesian_product.
func generateVariations(length int) []string {
	if length == 0 {
		return []string{""}
	}

	smallerCombinations := generateVariations(length - 1)
	var results []string

	// Add '+' and '*' to each combination
	for _, combination := range smallerCombinations {
		results = append(results, combination+"+")
		results = append(results, combination+"*")
	}

	return results
}

// calculateResults returns the totals of all equation which are valid.
// Valid is defined by whether the operands can be joined by +/* operators from left to right to reach the target.
func calculateResults(eq []equation) int {
	totalValid := 0

	for _, e := range eq {
		for _, o := range e.operators {

			total := e.operands[0]
			operators := strings.Split(o, "")

			for idx, operator := range operators {
				if operator == "+" {
					total += e.operands[idx+1]
				} else if operator == "*" {
					total *= e.operands[idx+1]
				}

				// If at the end of the current operator list, check if the running total matches the target
				if idx == len(operators)-1 && total == e.target {
					e.valid = true
					break
				}
			}
		}

		if e.valid {
			totalValid += e.target
		}
	}

	return totalValid
}
