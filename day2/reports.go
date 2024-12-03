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
	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("reading input file: %v", err)
	}

	reports, err := parseInput(inputBytes)
	if err != nil {
		log.Fatalf("parsing input: %v", err)
	}

	numberSafe, err := calculateNumSafeReports(reports)
	if err != nil {
		log.Fatalf("calculating the number of safe reports: %v", err)
	}

	fmt.Printf("Number of reports safe: %d (out of %d)\n", numberSafe, len(reports))
}

// parseInput parses a space separated slice of byte as a slice of int slices.
// See ./input.txt for example input.
func parseInput(input []byte) ([][]int, error) {
	rawReports := strings.Split(string(input), "\n")

	reports := make([][]int, 0, len(rawReports))

	for indexReport, report := range rawReports {
		levelsRaw := strings.Split(strings.TrimSpace(report), " ")

		levelsParsed := make([]int, 0, len(levelsRaw))

		for _, level := range levelsRaw {
			levelNumber, err := strconv.Atoi(level)
			if err != nil {
				return reports, fmt.Errorf("unable to convert string '%s' to an int on line %d: %v", level, indexReport, err)
			}

			levelsParsed = append(levelsParsed, levelNumber)
		}

		reports = append(reports, levelsParsed)
	}

	return reports, nil
}

// calculateNumSafeReports returns the number of safe reports in a slice of reports.
func calculateNumSafeReports(reports [][]int) (int, error) {
	numberSafe := 0

	for i, report := range reports {
		fmt.Printf("\nReport %d: %v\n", i, report)
		if checkReport(report) {
			numberSafe++
		}
	}

	return numberSafe, nil
}

// checkReport checks an individual report and returns whether it is safe or not.
// See readme.md for the rule logic.
func checkReport(report []int) bool {
	safe := true
	lastIndex := len(report) - 1

	// Check if the first two elements are increasing or decreasing
	increasing := true
	if report[0] > report[1] {
		increasing = false
	}

	for i := range report {
		if i == lastIndex {
			break
		}

		// Check if differing from the gradient (increasing vs decreasing)
		if increasing && (report[i] > report[i+1]) {
			fmt.Printf("Unsafe. Gradient changed. Switched from increasing to decreasing!\n")
			safe = false
		}

		if !increasing && (report[i] < report[i+1]) {
			fmt.Printf("Unsafe. Gradient changed. Switched from decreasing to increasing!\n")
			safe = false
		}

		// Check if within the 1 -> 3 difference range
		difference := absoluteNumber(report[i] - report[i+1])
		if difference < 1 || difference > 3 {
			fmt.Printf("Unsafe. Difference is greater than 3 or less than 1!\n")
			safe = false
		}
	}

	fmt.Printf("Reporting: %v\n", safe)
	return safe
}

func absoluteNumber(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
