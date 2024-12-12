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

/*

operator[0] = +
operator[1] = *

operand[0] = 81
operand[1] = 40
operand[2] = 27

Iterate over the operands?
operand[0]

*/

func main() {
	equations, err := parseConfig(inputFile)
	if err != nil {
		log.Fatalf("parsing config from %s: %v", inputFile, err)
	}

	totalValid := 0

	for _, e := range equations {
		e.operators = generateVariations(len(e.operands) - 1)

		for _, o := range e.operators {
			//fmt.Printf("original operators: %s\n", o)
			total := e.operands[0]
			operators := strings.Split(o, "")
			for idx, operator := range operators {
				//fmt.Printf("current operator: %s\n", operator)

				if operator == "+" {
					//fmt.Printf("operator is a +: adding %d to %d\n", e.operands[idx+1], total)
					total += e.operands[idx+1]

				} else if operator == "*" {
					//fmt.Printf("operator is a *: multiplying %d with %d\n", e.operands[idx+1], total)
					total *= e.operands[idx+1]

				} else {
					log.Fatalf("expected operator + or *, but got %s", operator)
				}

				//fmt.Printf("total=%d, target=%d\n", total, e.target)
				if total == e.target {
					e.valid = true
					break
				}
			}

			//fmt.Println()
		}

		if e.valid {
			totalValid += e.target
		}

		fmt.Printf("%#v\n", e)
	}

	fmt.Printf("Total = %d\n", totalValid)
}

type equation struct {
	target    int
	operands  []int
	operators []string
	valid     bool
}

func parseConfig(inputFilePath string) ([]equation, error) {
	b, err := os.ReadFile(inputFilePath)
	if err != nil {
		return []equation{}, fmt.Errorf("reading file from %s: %v", inputFile, err)
	}

	rows := strings.Split(string(b), "\n")
	results := make([]equation, 0, len(rows))

	for _, row := range rows {
		e := equation{
			operands:  make([]int, 0),
			operators: make([]string, 0),
		}

		rowSplit := strings.Split(row, ":")

		target, err := strconv.Atoi(rowSplit[0])
		if err != nil {
			return []equation{}, fmt.Errorf("converting string to int: %w", err)
		}
		e.target = target

		operands := strings.Split(strings.TrimSpace(rowSplit[1]), " ")
		for _, operand := range operands {
			o, err := strconv.Atoi(operand)
			if err != nil {
				return []equation{}, fmt.Errorf("converting string to int: %w", err)
			}
			e.operands = append(e.operands, o)
		}

		results = append(results, e)
	}

	return results, nil
}

// generateVariations generates all combinations of '+' and '*' for a given length
func generateVariations(length int) []string {
	if length == 0 {
		//fmt.Printf("Returning base case 0\n")
		return []string{""} // Base case: Return an empty string for length 0
	}

	// Recursive step: Generate combinations for (length - 1)
	//fmt.Printf("Calling generateVariations(%d)\n", length-1)
	smallerCombinations := generateVariations(length - 1)
	var results []string

	// Add '+' and '*' to each combination
	for _, combination := range smallerCombinations {
		//fmt.Printf("Appending %s to %s' (length=%d)\n", combination+"+", results, length)
		results = append(results, combination+"+")
		//fmt.Printf("Appending %s to %s' (length=%d)\n", combination+"*", results, length)
		results = append(results, combination+"*")
	}

	//fmt.Printf("Returning results: %v\n", results)
	return results
}
