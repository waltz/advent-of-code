package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("input-1-test.txt")
	if err != nil {
		fmt.Println("There was a problem reading the file.")
		return
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	matcher := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	for scanner.Scan() {
		currentLine := scanner.Text()
		matches := matcher.FindAllStringSubmatch(currentLine, -1)

		fmt.Println("current line", currentLine)

		for _, match := range matches {
			fmt.Println("got match", match)

			firstNumber, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println("error parsing the first number", match[1])
				return
			}

			secondNumber, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Println("error parsing the first number", match[2])
				return
			}

			multiplied := firstNumber * secondNumber
			total += multiplied

			fmt.Println("multiplied:", multiplied)
		}

		fmt.Println("line processed, total:", total)
	}

	fmt.Println("finished, multiplication total:", total)
}
