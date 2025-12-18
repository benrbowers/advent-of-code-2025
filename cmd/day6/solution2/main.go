package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
---- Part 1 ----
- A young cephalopod needs help with her math homework.
- The input is in the form of:
	- Columns of numbers
	- Followed by either a "+" or a "*"
- The answer for each column is the result of the numbers
	being combined using the operator below.
- To find the answer:
	- Calculate the sum of the results for every column.

---- Part 2 ----
- The big cephalopods return to inform you the answer is wrong.
- Cephalopod math works differently:
	- Numbers are written VERTICALLY, downwards
	- And from RIGHT-TO-LEFT
		- Because operators are strictly + and *,
			the order doesn't actually matter.
*/

func main() {
	operatorIndexes := []int{}
	numberLines := []string{}
	var operatorLine string
	var resultSum int = 0

	input, err := os.Open("./cmd/day6/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		if line[0] == '+' || line[0] == '*' {
			operatorLine = line
		} else {
			numberLines = append(numberLines, line)
		}
	}

	for i, char := range operatorLine {
		if char == '+' || char == '*' {
			operatorIndexes = append(operatorIndexes, i)
		}
	}

	numberRows := [][]int{}
	for i, op := range operatorIndexes {
		var next int
		if i < len(operatorIndexes)-1 {
			next = operatorIndexes[i+1]
		} else {
			next = len(operatorLine) + 1
		}

		for col := op; col < next-1; col++ {
			numberString := []byte{}

			for _, line := range numberLines {
				if line[col] != ' ' {
					numberString = append(numberString, line[col])
				}
			}

			number, err := strconv.Atoi(string(numberString))
			if err != nil {
				panic(err)
			}

			if i >= len(numberRows) {
				numberRows = append(numberRows, []int{number})
			} else {
				numberRows[i] = append(numberRows[i], number)
			}
		}
	}

	for i, row := range numberRows {
		var result int

		operator := operatorLine[operatorIndexes[i]]

		if operator == '+' {
			result = 0
		} else {
			result = 1
		}

		for _, number := range row {
			if operator == '+' {
				result += number
			} else {
				result *= number
			}
		}

		resultSum += result
	}

	fmt.Println("Sum of every column's answer:", resultSum)
}
