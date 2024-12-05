package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("input-1-test.txt")
	if err != nil {
		fmt.Println("There was a problem reading the file.")
		return
	}

	safeReports := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rawReport := strings.Split(scanner.Text(), " ")

		report := []int{}

		for _, rawLevel := range rawReport {
			level, err := strconv.Atoi(rawLevel)

			if err != nil {
				fmt.Println("error parsing a level", level)
				return
			}

			report = append(report, level)
		}

		isSafe := true
		isIncreasing := true

		for i := 0; i < (len(report) - 1); i++ {
			difference := report[i] - report[i+1]

			// establish the trend on our first pass
			if i == 0 && difference > 0 {
				isIncreasing = false
			}

			// on successive passes, make sure we continue with the trend.
			if i > 0 && difference > 0 && isIncreasing == true {
				// fmt.Println("the difference is decreasing, but it should be increasing. unsafe.", difference, report[i], report[i+1])
				isSafe = false
				break
			}

			// same here, but if the trend is in the opposite direction
			if i > 0 && difference < 0 && isIncreasing == false {
				// fmt.Println("the difference is increasing, but it should be decreasing. unsafe.", difference, report[i], report[i+1])
				isSafe = false
				break
			}

			// make sure that the difference is never greater than three or less than 1
			if difference < 0 {
				difference = -difference
			}

			// fmt.Println(difference)
			if difference < 1 || difference > 3 {
				// fmt.Println("the difference was less than one or greater than three", difference)
				isSafe = false
				break
			}
		}

		fmt.Println("report safety", report, isSafe)

		if isSafe == true {
			safeReports += 1
		}
	}

	fmt.Println("finished, safe reports count:", safeReports)
}
