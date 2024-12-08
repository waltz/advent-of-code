package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("input-2-test.txt")
	if err != nil {
		fmt.Println("There was a problem reading the file.")
		return
	}
	defer file.Close()

	total := 0
	multiplicationEnabled := true
	scanner := bufio.NewScanner(file)
	matcher := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don\'t\(\)`)

	for scanner.Scan() {
		currentLine := scanner.Text()
		matches := matcher.FindAllStringSubmatch(currentLine, -1)

		fmt.Println("current line", currentLine)

		for _, match := range matches {
			instruction := match[0]

			fmt.Println("processing instruction:", instruction)

			if strings.Contains(instruction, "mul") && multiplicationEnabled == true {
				// if multiplicationEnabled == false {
				// 	fmt.Println("skipping multiplication, currently disabled")
				// 	break
				// }

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

				fmt.Println("multiplied", firstNumber, secondNumber, multiplied)
			} else if strings.Contains(instruction, "don't") {
				multiplicationEnabled = false
				fmt.Println("disabling multiplication")
			} else if strings.Contains(instruction, "do") {
				multiplicationEnabled = true
				fmt.Println("enabling multiplication")
			} else {
				fmt.Println("found and unrecognized instruction:", instruction)
			}
		}

		fmt.Println("line processed, total:", total)
	}

	fmt.Println("finished, multiplication total:", total)
}
