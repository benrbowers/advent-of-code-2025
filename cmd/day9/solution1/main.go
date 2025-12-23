package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
- The elves in the movie theatre are decorating by
	rearranging the floor tiles.
- Some of the tiles are red.
- The elves would like to know:
	- What is the LARGEST rectangle that uses two red
		tiles as its corners?
- The input is a list of coords of red tiles.
- To find the answer:
	- What is the area of the largest rectangle you
		can create?
*/

func main() {
	tiles := [][2]int{}
	var largestArea int = 0

	input, err := os.Open("./cmd/day9/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")

		if len(parts) != 2 {
			panic(`Invalid tile coord: "` + scanner.Text() + `"`)
		}

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(fmt.Errorf("Failed to parse X coord: %w", err))
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(fmt.Errorf("Failed to parse Y coord: %w", err))
		}

		tiles = append(tiles, [2]int{x, y})
	}

	for i := 0; i < len(tiles)-1; i++ {
		for j := i + 1; j < len(tiles); j++ {
			dx := abs(tiles[j][0]-tiles[i][0]) + 1
			dy := abs(tiles[j][1]-tiles[i][1]) + 1
			area := dx * dy

			if area > largestArea {
				largestArea = area
			}
		}
	}

	fmt.Println("Largest possible area with red corners: ", largestArea)
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
