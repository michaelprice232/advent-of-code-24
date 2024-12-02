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
		log.Fatalf("Error reading input file: %v", err)
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

func parseInput(input []byte) ([][]int, error) {
	rawReports := strings.Split(string(input), "\n")

	reports := make([][]int, 0)

	for indexReport, report := range rawReports {
		levelsRaw := strings.Split(strings.TrimSpace(report), " ")

		levelsParsed := make([]int, 0)

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

func calculateNumSafeReports(reports [][]int) (int, error) {
	numberSafe := 0

	for i, report := range reports {
		fmt.Printf("\nReport %d\n", i)
		fmt.Println(report)
		if checkReport(report) {
			numberSafe++
		}
	}

	return numberSafe, nil
}

func checkReport(report []int) bool {
	safe := true
	lastIndex := len(report) - 1
	//fmt.Printf("Last index: %d\n", lastIndex)

	// Check if the first two elements are increasing or decreasing
	increasing := true
	if report[0] > report[1] {
		increasing = false
	}

	for i := range report {
		if i != lastIndex {
			// Check if differing from the gradient
			if increasing && (report[i] > report[i+1]) {
				fmt.Printf("Reporting as false: gradient changed!\n")
				safe = false
				break
			}

			if !increasing && (report[i] < report[i+1]) {
				fmt.Printf("Reporting as false: gradient changed!\n")
				safe = false
				break
			}

			// Check if within the 1 -> 3 difference
			if absoluteNumber(report[i]-report[i+1]) > 3 || absoluteNumber(report[i]-report[i+1]) < 1 {
				fmt.Printf("Reporting as false: difference is greater than 3 or less than 1!\n")
				safe = false
				break
			}
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
