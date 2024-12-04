package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	first_list := []int{}
	second_list := []int{}

	// read the file
	file, err := os.Open("input.txt")
	// file, err := os.Open("input-test.txt")
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
		first_list = append(first_list, first_number)

		second_number, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println("error converting the second number", err)
			return
		}
		second_list = append(second_list, second_number)

		if err := scanner.Err(); err != nil {
			fmt.Println("there was an issue reading a line")
			return
		}
	}

	firstSlice := sort.IntSlice(first_list)
	secondSlice := sort.IntSlice(second_list)

	firstSlice.Sort()
	secondSlice.Sort()

	fmt.Println("sorted and deduped arrays", len(first_list), len(second_list))

	var finalDistance int = 0

	for index, firstNumber := range first_list {
		var secondIndex = 0

		if index > (len(second_list) - 1) {
			secondIndex = len(second_list) - 1
		} else {
			secondIndex = index
		}

		var secondNumber = second_list[secondIndex]
		var distance = firstNumber - secondNumber

		if distance < 0 {
			distance = -distance
		}

		finalDistance = finalDistance + distance

		fmt.Println("got numbers", index, firstNumber, secondIndex, secondNumber, distance)
	}

	fmt.Println("finished", finalDistance)
}
