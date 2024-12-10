package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*

WIP!!!

Generate all the possible combinations of operators based on an arbitrary number of operands:
https://en.wikipedia.org/wiki/Cartesian_product
*/

const inputFile = "./input.txt"

func main() {
	equ, err := parseConfig(inputFile)
	if err != nil {
		log.Fatalf("parsing config from %s: %v", inputFile, err)
	}

	for _, e := range equ {
		e.operators = generateVariations(len(e.operands) - 1)

		fmt.Printf("%#v\n", e)

		for i, operand := range e.operands {

		}
	}

	//// Example: Generate all combinations for length 3
	//length := 3
	//fmt.Printf("Calling generateVariations(%d)\n", length)
	//results := generateVariations(length)
	//
	//// Print all results
	//for _, variation := range results {
	//	fmt.Println(variation)
	//}
}

type equations struct {
	target    int
	operands  []int
	operators []string
}

func parseConfig(inputFilePath string) ([]equations, error) {
	b, err := os.ReadFile(inputFilePath)
	if err != nil {
		return []equations{}, fmt.Errorf("reading file from %s: %v", inputFile, err)
	}

	rows := strings.Split(string(b), "\n")
	results := make([]equations, 0, len(rows))

	for _, row := range rows {
		equation := equations{
			operands:  make([]int, 0),
			operators: make([]string, 0),
		}

		rowSplit := strings.Split(row, ":")

		target, err := strconv.Atoi(rowSplit[0])
		if err != nil {
			return []equations{}, fmt.Errorf("converting string to int: %w", err)
		}
		equation.target = target

		operands := strings.Split(strings.TrimSpace(rowSplit[1]), " ")
		for _, operand := range operands {
			o, err := strconv.Atoi(operand)
			if err != nil {
				return []equations{}, fmt.Errorf("converting string to int: %w", err)
			}
			equation.operands = append(equation.operands, o)
		}

		results = append(results, equation)
	}

	return results, nil
}

// generateVariations generates all combinations of '+' and '*' for a given length
func generateVariations(length int) []string {
	if length == 0 {
		fmt.Printf("Returning base case 0\n")
		return []string{""} // Base case: Return an empty string for length 0
	}

	// Recursive step: Generate combinations for (length - 1)
	fmt.Printf("Calling generateVariations(%d)\n", length-1)
	smallerCombinations := generateVariations(length - 1)
	var results []string

	// Add '+' and '*' to each combination
	for _, combination := range smallerCombinations {
		fmt.Printf("Appending %s to %s' (length=%d)\n", combination+"+", results, length)
		results = append(results, combination+"+")
		fmt.Printf("Appending %s to %s' (length=%d)\n", combination+"*", results, length)
		results = append(results, combination+"*")
	}

	fmt.Printf("Returning results: %v\n", results)
	return results
}
