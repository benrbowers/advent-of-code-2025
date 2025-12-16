package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
---- Part 1 ----
- The elves in the cafeteria can't figure out which
	ingredients are spoiled.
- Input is in the form of:
	- Ranges (x-y) representing which ID's are FRESH
	- A blank line
	- ID's for the available ingredients
- To find the answer:
	- How many available ingredients are fresh?

---- Part 2 ----
- Now, the elves want to know EVERY ID that falls
	within a fresh range.
- The 2nd section of the input is now IRRELEVANT.
- To find the answer:
	- How many UNIQUE id's fall within fresh ranges?
*/

func main() {
	freshIDs := []int{}

	input, err := os.Open("./cmd/day5/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

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

		for id := start; id <= stop; id++ {
			if !slices.Contains(freshIDs, id) {
				// ISSUE: freshIDs gets so big that slice.Contains
				// takes a long time to execute. And it must be
				// executed for EVERY id in the range.
				// Likely solution: Use simple addition and subtraction
				freshIDs = append(freshIDs, id)
			}
		}
	}

	fmt.Println("Number of fresh ID's:", len(freshIDs))
}
