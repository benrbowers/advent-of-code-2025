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
	uniqueRanges := [][2]int{}

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

		// Merge overlapping ranges

		var insertIndex int = -1
		var removeIndex int = -1

		for i, uniqueRange := range uniqueRanges {
			if (start >= uniqueRange[0] && start <= uniqueRange[1]) ||
				(stop >= uniqueRange[0] && stop <= uniqueRange[1]) ||
				(uniqueRange[0] >= start && uniqueRange[0] <= stop) ||
				(uniqueRange[1] >= start && uniqueRange[1] <= stop) {
				// Ranges overlap
				start = min(start, uniqueRange[0])
				stop = max(stop, uniqueRange[1])

				if insertIndex == -1 {
					insertIndex = i
				}
				removeIndex = i + 1
			} else if stop < uniqueRange[0] {
				// stop is before start
				if insertIndex == -1 {
					insertIndex = i
				}
				break
			}
		}

		if removeIndex > -1 {
			uniqueRanges = slices.Delete(
				uniqueRanges,
				insertIndex,
				removeIndex,
			)
		}

		if insertIndex > -1 {
			uniqueRanges = slices.Insert(
				uniqueRanges,
				insertIndex,
				[2]int{start, stop},
			)
		} else {
			uniqueRanges = append(uniqueRanges, [2]int{start, stop})
		}
	}

	var numFreshIDs int = 0
	for _, uniqueRange := range uniqueRanges {
		numFreshIDs += uniqueRange[1] - uniqueRange[0] + 1
	}

	fmt.Println("Number of fresh ID's:", numFreshIDs)
}
