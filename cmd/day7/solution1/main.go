package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
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
*/

func main() {
	prevBeamRow := []int{}
	var splitCount int = 0

	input, err := os.Open("./cmd/day7/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)

	var row int = 0

	for scanner.Scan() {
		row++
		if row%2 == 0 {
			// Since beams are vertical, we can skip empty lines
			continue
		}

		newBeamRow := []int{}
		line := scanner.Text()

		if row == 1 {
			sourceIndex := strings.Index(line, "S")
			newBeamRow = append(newBeamRow, sourceIndex)
		} else {
			for _, beam := range prevBeamRow {
				if line[beam] == '^' {
					splitCount++
					if len(newBeamRow) == 0 || newBeamRow[len(newBeamRow)-1] != beam-1 {
						newBeamRow = append(newBeamRow, beam-1)
					}
					newBeamRow = append(newBeamRow, beam+1)
				} else {
					// Propagate down empty space
					if len(newBeamRow) == 0 || newBeamRow[len(newBeamRow)-1] != beam {
						newBeamRow = append(newBeamRow, beam)
					}
				}
			}
		}

		prevBeamRow = newBeamRow
	}

	fmt.Println("Number of times beam split:", splitCount)
}
