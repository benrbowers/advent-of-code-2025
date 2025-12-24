package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
---- Part 1 ----
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

---- Part 2 ----
- Green tiles now connect all the red tiles, in the
	order they are listed.
- The new shape that is formed is FILLED with green tiles.
- To find the answer:
	- What is the area of the largest rectangle that has
		TWO red corners and ONLY red/green tiles?
*/

// Checking all green squares for each square in rect
// is very expensive.
// Cheaper: For all EMPTY SQUARES, check if in rect's
// range

// 1. Find shape boundary
// 2. Collect all non-green/red tiles
// 3. Collect all pairs, order by area desc
// 4. For largest pairs, which has no non-green/red tiles

type Rectangle struct {
	corners [2]int // indexes of `redTiles`
	area    int
}

func main() {
	redTiles := [][2]int{}
	boundary := [][2]int{}
	invalidTiles := [][2]int{}
	var minX int = math.MaxInt
	var maxX int = 0
	var minY int = math.MaxInt
	var maxY int = 0
	redPairs := []Rectangle{}
	var largestGreenRedRectangle Rectangle

	input, err := os.Open("./cmd/day9/test.txt")
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

		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}

		redTiles = append(redTiles, [2]int{x, y})
	}

	// Find boundary
	for i, red := range redTiles {
		boundary = append(boundary, red)

		var next [2]int
		if i == len(redTiles)-1 {
			// Connect back to first
			next = redTiles[0]
		} else {
			next = redTiles[i+1]
		}

		if red[0] == next[0] {
			start := red[1] + 1
			stop := next[1]
			if red[1] > next[1] {
				start = next[1] + 1
				stop = red[1]
			}

			x := red[0]

			for y := start; y < stop; y++ {
				boundary = append(boundary, [2]int{x, y})
			}
		} else {
			start := red[0] + 1
			stop := next[0]
			if red[0] > next[0] {
				start = next[0] + 1
				stop = red[0]
			}

			y := red[1]

			for x := start; x < stop; x++ {
				boundary = append(boundary, [2]int{x, y})
			}
		}
	}

	// Collect invalid tiles within shape rect
	// (inside the rect, OUTSIDE the green shape)
	// BOTTLENECK: Collecting EVERY invalid square takes too long
	// Instead, for each large rect, check if boundary is contained in rect range
	// N of boundary is MUCH smaller than N (area) of invalid tiles.

	for x := minX; x <= maxX; x++ {
		var insideBoundary bool = false
		var insideShape bool = false
		for y := minY; y <= maxY; y++ {
			if slices.Contains(boundary, [2]int{x, y}) {
				insideBoundary = true
				continue
			}

			if insideBoundary {
				if x != minX && maxX != x &&
					y != minY && y != maxY {
					insideShape = !insideShape
				}
				insideBoundary = false
			}

			if !insideShape {
				invalidTiles = append(invalidTiles, [2]int{x, y})
			}
		}
	}

	for i := 0; i < len(redTiles)-1; i++ {
		for j := i + 1; j < len(redTiles); j++ {
			dx := abs(redTiles[j][0]-redTiles[i][0]) + 1
			dy := abs(redTiles[j][1]-redTiles[i][1]) + 1
			area := dx * dy

			redPairs = append(redPairs, Rectangle{
				[2]int{i, j},
				area,
			})
		}
	}

	// Sort descending by area
	slices.SortFunc(redPairs, func(a, b Rectangle) int {
		return b.area - a.area
	})

	// Find first pair with NO invalid tiles
PairLoop:
	for _, pair := range redPairs {
		c1 := redTiles[pair.corners[0]]
		c2 := redTiles[pair.corners[1]]

		for _, tile := range invalidTiles {
			minX := min(c1[0], c2[0])
			maxX := max(c1[0], c2[0])
			minY := min(c1[1], c2[1])
			maxY := max(c1[1], c2[1])

			if tile[0] >= minX && tile[0] <= maxX &&
				tile[1] >= minY && tile[1] <= maxY {
				// Rectangle contains invalid tile
				continue PairLoop
			}
		}

		largestGreenRedRectangle = pair
		break
	}

	fmt.Println(
		"Largest possible green/red area with red corners:",
		largestGreenRedRectangle.area,
	)
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
