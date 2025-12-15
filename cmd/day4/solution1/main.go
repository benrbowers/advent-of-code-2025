package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
- You need to optimize the work of forklifts lifting rolls of paper.
- Input is a grid of:
	- empty spaces: "."
	- paper rolls: "@"
- A roll of paper can only be accessed if there are
	FEWER THAN FOUR rolls in the adjacent eight spaces.
*/

func main() {
	rollGrid := [][]bool{} // bools representing whether a roll is present
	var accessibleCount int = 0

	input, err := os.Open("./cmd/day4/input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var row int = 0

	for scanner.Scan() {
		line := scanner.Text()

		rollGrid = append(rollGrid, make([]bool, len(line)))

		for col, slot := range line {
			rollGrid[row][col] = slot == '@'
		}

		row++
	}

	height := len(rollGrid)
	width := len(rollGrid[0])

	for row = range height {
		for col := range width {
			if rollGrid[row][col] {
				var rollCount int = 0

				// Loop "around" the current pos
				for y := -1; y <= 1; y++ {
					if row+y < 0 || row+y >= height {
						continue
					}
					for x := -1; x <= 1; x++ {
						if col+x < 0 || col+x >= width {
							continue
						}
						if x == 0 && y == 0 {
							continue
						}

						if rollGrid[row+y][col+x] {
							rollCount++
						}
					}
				}

				if rollCount < 4 {
					accessibleCount++
				}
			}
		}
	}

	fmt.Println("Number of accessible rolls:", accessibleCount)
}
