package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
- A young cephalopod needs help with her math homework.
- The input is in the form of:
	- Columns of numbers
	- Followed by either a "+" or a "*"
- The answer for each column is the result of the numbers
	being combined using the operator below.
- To find the answer:
	- Calculate the sum of the results for every column.
*/

func main() {
	numberColumns := [][]int{}
	operators := []string{}
	var resultSum int = 0

	input, err := os.Open("./cmd/day6/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " ")

		if splitLine[0] == "+" || splitLine[0] == "*" {
			for _, part := range splitLine {
				if part == " " || part == "" {
					continue
				}

				operators = append(operators, part)
			}
		} else {
			var columnIndex = 0

			for _, part := range splitLine {
				if part == " " || part == "" {
					continue
				}

				number, err := strconv.Atoi(part)
				if err != nil {
					panic(err)
				}

				if columnIndex >= len(numberColumns) {
					numberColumns = append(numberColumns, []int{number})
				} else {
					numberColumns[columnIndex] = append(
						numberColumns[columnIndex],
						number,
					)
				}

				columnIndex++
			}
		}
	}

	for i, column := range numberColumns {
		var result int

		if operators[i] == "+" {
			result = 0
		} else {
			result = 1
		}

		for _, number := range column {
			if operators[i] == "+" {
				result += number
			} else {
				result *= number
			}
		}

		resultSum += result
	}

	fmt.Println("Sum of every column's answer:", resultSum)
}
