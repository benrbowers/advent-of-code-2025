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
- The elves in the playground are decorating by
	haning up electrical junction boxes.
- When lights are connected to TWO junction boxes,
	electicity can flow between them.
- The elves would like:
	- To connect the junction boxes that are as CLOSE
		TOGETHER AS POSSIBLE, to save lights.
- Input is in the form of:
	- A list of 3D coords representing the positions
		of the junction boxes.
- To find the answer:
	- Connect together the CLOSEST 1000 pairs of
		junction boxes
	- Multiply together the sizes of the THREE LARGEST
		circuits
*/

// Calc all pairs: n(n-1), 999000 for n = 1000
// Brute force very doable for ~1M

type Pair struct {
	boxes [2]int
	dist  float64
}

func main() {
	circuits := [][]int{}
	boxes := [][3]int{}

	// Getting correct answer for test, but not input.
	// Not sure where to go from here.
	// Need better test cases.
	input, err := os.Open("./cmd/day8/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")
		if len(coords) != 3 {
			panic(`Invalid coord format: "` + line + `"`)
		}

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			panic(fmt.Errorf("Failed to parse X coord: %w", err))
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			panic(fmt.Errorf("Failed to parse Y coord: %w", err))
		}
		z, err := strconv.Atoi(coords[2])
		if err != nil {
			panic(fmt.Errorf("Failed to parse Z coord: %w", err))
		}

		boxes = append(boxes, [3]int{x, y, z})
	}

	pairs := []Pair{}

	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			pairs = append(pairs, Pair{
				[2]int{i, j},
				dist3D(boxes[i], boxes[j]),
			})
		}
	}

	slices.SortFunc(pairs, func(a, b Pair) int {
		if a.dist-b.dist < 0.0 {
			return -1
		} else {
			return 1
		}
	})

	var pairIndex int = 0
	var connectionCount int = 0
ConnectionLoop:
	for connectionCount < 1000 {
		p := pairs[pairIndex]
		pairIndex++

		for i, circuit := range circuits {
			has0 := slices.Contains(circuit, p.boxes[0])
			has1 := slices.Contains(circuit, p.boxes[1])

			if has0 && has1 {
				// Pair is already connected, skip
				continue ConnectionLoop
			}

			if has0 {
				circuits[i] = append(circuits[i], p.boxes[1])
			} else if has1 {
				circuits[i] = append(circuits[i], p.boxes[0])
			}

			if has0 || has1 {
				connectionCount++
				continue ConnectionLoop
			}
		}

		circuits = append(circuits, []int{
			p.boxes[0], p.boxes[1],
		})
		connectionCount++
	}

	slices.SortFunc(circuits, func(a, b []int) int {
		return len(b) - len(a)
	})

	var result int = 1
	for i := range 3 {
		result *= len(circuits[i])
	}
	fmt.Println("Sizes of the three largest circuits multiplied together: ", result)
	fmt.Println("Total circuits:", len(circuits))
}

func dist3D(start, stop [3]int) float64 {
	difX := math.Pow(float64(start[0]-stop[0]), 2)
	difY := math.Pow(float64(start[1]-stop[1]), 2)
	difZ := math.Pow(float64(start[2]-stop[2]), 2)

	return math.Sqrt(difX + difY + difZ)
}
