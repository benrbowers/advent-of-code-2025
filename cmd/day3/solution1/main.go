package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
- You need to restore power to the escalator using emergency batteries.
- Input is in the form of single digit numbers 1-9 representing
  the "joltage" level of all the batteries.
- Battieries are organized into "Banks," represented
  by a single line of the input.
- Within each bank, you need to turn on EXACTLY TWO batteries.
- The digits of the two turned on batteries form the bank's joltage level.
- You need to find the maximum possible joltage each bank can produce.
- To find the answer: Calculate the sum of the max joltage levels of each bank.
*/

// Cases to consider:
// 111121119112 -> 92
// 1111911121119 -> 92
// 1111111211191 -> 91

func main() {
	input, err := os.Open("./cmd/day3/input.txt")
	if err != nil {
		panic(err)
	}

	var totalOutputJoltage int = 0

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		digits := scanner.Bytes()

		var joltMax1 byte = 0
		var maxPos1 int = 0
		var joltMax2 byte = 0
		var maxPos2 int = 0

		for pos, digit := range digits {
			if digit > joltMax1 {
				// Use first occurrance of largest digit
				joltMax2 = joltMax1
				joltMax1 = digit
				maxPos2 = maxPos1
				maxPos1 = pos
			} else if digit >= joltMax2 {
				// Use LAST occurrance of second largest digit
				joltMax2 = digit
				maxPos2 = pos
			}
		}

		var joltString string

		if maxPos1 < maxPos2 {
			joltString = string([]byte{joltMax1, joltMax2})
		} else {
			// Check if room after maxPos1, find largest after
			// To cover: 1111111211191 -> 91
			if maxPos1 < len(digits)-1 {
				var nextMax byte = 0

				for i := maxPos1 + 1; i < len(digits); i++ {
					if digits[i] > nextMax {
						nextMax = digits[i]
					}
				}

				joltString = string([]byte{joltMax1, nextMax})
			} else {
				joltString = string([]byte{joltMax2, joltMax1})
			}
		}

		bankJoltage, err := strconv.Atoi(joltString)
		if err != nil {
			panic(fmt.Errorf("Failed to parse bankJoltage: %w", err))
		}

		totalOutputJoltage += bankJoltage
	}

	fmt.Println("Total output joltage:", totalOutputJoltage)
}
