package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const inputPath = "./input.txt"

func main() {
	bytes, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("reading file %s: %v", inputPath, err)
	}

	listOne, listTwo, err := processLists(bytes)
	if err != nil {
		log.Fatalf("processing file %s: %v", inputPath, err)
	}

	displayResults(listOne, listTwo)
}

// processLists sorts two int slices in order. It also confirms there are 2 numbers per line, and they are equal length.
func processLists(input []byte) ([]int, []int, error) {
	listOne := make([]int, 0)
	listTwo := make([]int, 0)

	for i, line := range strings.Split(string(input), "\n") {
		linePairs := strings.Split(line, "   ")
		if len(linePairs) != 2 {
			return listOne, listTwo, fmt.Errorf("expected 2 numbers per line. got %d on line %d", len(linePairs), i)
		}

		result1, err := strconv.Atoi(linePairs[0])
		if err != nil {
			return listOne, listTwo, fmt.Errorf("parsing line %d: %v", i, err)
		}
		result2, err := strconv.Atoi(linePairs[1])
		if err != nil {
			return listOne, listTwo, fmt.Errorf("parsing line %d: %v", i, err)
		}

		listOne = append(listOne, result1)
		listTwo = append(listTwo, result2)
	}

	slices.Sort(listOne)
	slices.Sort(listTwo)

	if len(listOne) != len(listTwo) {
		log.Fatalf("lists are of different lengths: %d vs %d", len(listOne), len(listTwo))
	}

	return listOne, listTwo, nil
}

// calculateDistance calculates the distance between two integers.
func calculateDistance(x, y int) int {
	return absoluteNumber(x - y)
}

// absoluteNumber returns the absolute number.
func absoluteNumber(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// displayResults prints out the results to the terminal.
func displayResults(listOne []int, listTwo []int) {
	log.Printf("Smallest numbers: %d (list 1) vs %d (list 2)", listOne[0], listTwo[0])
	log.Printf("Largest numbers: %d (list 1) vs %d (list 2)", listOne[len(listOne)-1], listTwo[len(listTwo)-1])

	total := 0
	for i, v := range listOne {
		total += calculateDistance(v, listTwo[i])
	}

	log.Printf("Total distance: %d", total)
}
