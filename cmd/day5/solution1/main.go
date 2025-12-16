package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
- The elves in the cafeteria can't figure out which
	ingredients are spoiled.
- Input is in the form of:
	- Ranges (x-y) representing which ID's are FRESH
	- A blank line
	- ID's for the available ingredients
- To find the answer:
	- How many available ingredients are fresh?
*/

func main() {
	var rangesNotParsed bool = true
	freshRanges := [][2]int{}
	var freshCount int = 0

	input, err := os.Open("./cmd/day5/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)

ScanLoop:
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			rangesNotParsed = false
			continue
		}

		if rangesNotParsed {
			startStop := strings.Split(line, "-")

			if len(startStop) != 2 {
				panic(`Invalid format for ID range: "` + line + `"`)
			}

			start, err := strconv.Atoi(startStop[0])
			if err != nil {
				panic(fmt.Errorf("Failed to parse ID range start: %w", err))
			}
			stop, err := strconv.Atoi(startStop[1])
			if err != nil {
				panic(fmt.Errorf("Failed to parse ID range stop: %w", err))
			}

			freshRanges = append(freshRanges, [2]int{start, stop})
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				panic(fmt.Errorf("Failed to parsed ingredient ID: %w", err))
			}

			for _, idRange := range freshRanges {
				if id >= idRange[0] && id <= idRange[1] {
					freshCount++
					continue ScanLoop
				}
			}
		}
	}

	fmt.Println("Number of fresh ID's:", freshCount)
}
