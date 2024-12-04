package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	firstList := []int{}
	secondList := []int{}

	// read the file
	// file, err := os.Open("input.txt")
	file, err := os.Open("input-2-test.txt")
	if err != nil {
		fmt.Println("There was a problem reading the file.")
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// split on three spaces, because the input file uses that as a delimiter.
		numbers := strings.Split(scanner.Text(), "   ")

		first_number, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println("error converting the first number")
			return
		}
		firstList = append(firstList, first_number)

		second_number, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println("error converting the second number", err)
			return
		}
		secondList = append(secondList, second_number)

		if err := scanner.Err(); err != nil {
			fmt.Println("there was an issue reading a line")
			return
		}
	}

	// create a map of numbers and occurrences from the second list
	secondListOccurrences := make(map[int]int)

	for _, number := range secondList {
		occurrencesForValue, exists := secondListOccurrences[number]

		if !exists {
			secondListOccurrences[number] = 1
		} else {
			secondListOccurrences[number] = occurrencesForValue + 1
		}
	}

	fmt.Println("occurrences", secondListOccurrences)

	var similarityScore = 0

	for _, number := range firstList {
		occurrences, exists := secondListOccurrences[number]

		var multiplier = 0

		if exists {
			multiplier = occurrences
		}

		similarityScore = similarityScore + (number * multiplier)
	}

	fmt.Println("finished", similarityScore)
}
