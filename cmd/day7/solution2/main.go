package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
---- Part 1 ----
- You need to repair the teleporter to get back to
	the research wing.
- The teleporter error code suggests an issue with
	"tachyon manifolds."
- The input is in the form of:
	- A capital "S" marking the initial source of the
		tachyon beam
	- Empty spaces represented by "."
	- Beam splitters represented by "^"
- When a beam is split, two new beams that are also
	FULLY VERTICAL form in the two spaces immediately
	adjacent to the beam splitter.
- To find the answer:
	- How many times will the initial beam be split?

---- Part 2 ----
- It turns out you are actually dealing with a
	"quantum" tachyon manifold
- With a quantum tachyon manifold, only a SINGLE tachyon
	particle is sent through the manifold.
- A tachyon particle takes BOTH the left and right path
	of each splitter encountered, in DIFFERENT TIMELINES.
- To find the answer:
	- How many different timelines would a single tachyon
		particle end up on?
*/

// Almost definitely going to need recursion.
// Base case: Tachyon exits manifold
// Base + 1: Tachyon splits into two timelines

var lines = []string{}

func main() {
	input, err := os.Open("./cmd/day7/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)

	var row int = 0

	for scanner.Scan() {
		row++
		if row%2 == 0 {
			continue
		}
		lines = append(lines, scanner.Text())
	}

	sourceIndex := strings.Index(lines[0], "S")

	timelineCount := countTimelines([2]int{1, sourceIndex})

	fmt.Println("Timelines encountered by tachyon:", timelineCount)
}

func countTimelines(start [2]int) int {
	// Recursion is correct because works on test input -> 40
	// But run time too large for large n
	// Possible to calc timelines "one at a time"?
	// Is memoization possible?
	// Possible with math if we assume every splitter gets hit?
	if start[0] == len(lines) {
		return 1
	}

	if lines[start[0]][start[1]] == '^' {
		return countTimelines(
			[2]int{start[0] + 1, start[1] - 1},
		) + countTimelines(
			[2]int{start[0] + 1, start[1] + 1},
		)
	}

	return countTimelines([2]int{start[0] + 1, start[1]})
}
