package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
---- PART 1 ----
- Trying to crack safe with numbers 0-99
- Input is a sequence of "rotations" that tell you how to open the safe.
- A rotation includes:
	- an "L" or an "R" to indicate left or right.
	- an integer "distance" representing how many clicks to rotate
- Dial is circular and returns to 0: 99 -> 0
- The dial STARTS at 50
- The safe is a DECOY
- The actual answer is:
	- the number of times the dial is left pointing at 0 after any rotation in the sequence.

	---- PART 2 ----
- The combination is still wrong.
- The REAL answer is the TOTAL number of 0's encountered
*/

func main() {
	var dial int = 50
	var zeroClicks int = 0

	input, err := os.Open("./cmd/day1/input.txt")
	if err != nil {
		panic(fmt.Errorf("Failed to open input file: %w", err))
	}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 2 {
			panic(fmt.Errorf(`Invalid rotation: "%s"`, line))
		}

		direction := line[0:1]

		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(fmt.Errorf(`Failed to parse distance: %w`, err))
		}

		zeroClicks += distance / 100 // int division for full rotations
		remainder := distance % 100
		switch direction {
		case "R":
			if dial+remainder > 99 {
				zeroClicks += 1
			}
		case "L":
			if dial-remainder <= 0 {
				zeroClicks += 1
			}
		}

		dial = turnDial(dial, direction, distance)
	}
	fmt.Println("0's encountered:", zeroClicks)
}

func turnDial(dial int, direction string, distance int) int {
	switch direction {
	case "R":
		return (dial + distance) % 100
	case "L":
		diff := dial - (distance % 100)
		if diff < 0 {
			return 100 + diff
		} else {
			return diff
		}
	default:
		return dial
	}
}
