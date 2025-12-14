package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
- Invalid ID's have been entered into the gift shop database by a young elf.
- The young elf entered the ID's by repeating numbers twice: e.g., 99, 6464, 123123
- Input is in the form of a single line of comma separated ranges, representing the ranges to check
- NONE of the ID's have LEADING ZEROES: 0101 does not exist, 101 is a valid ID, 1010 is an invalid ID
- To find the answer: add up all of the invalid ID's
*/

func main() {
	var invalidTotal int = 0

	input, err := os.Open("./cmd/day2/input.txt")
	if err != nil {
		panic(fmt.Errorf("Failed to open input file: %w", err))
	}
	defer input.Close()

	reader := bufio.NewReader(input)

	for {
		idRange, readErr := reader.ReadString(',')

		if readErr != nil && readErr != io.EOF {
			panic(readErr)
		}

		if idRange[len(idRange)-1] == ',' || idRange[len(idRange)-1] == '\n' {
			// Remove comma or new line if necessary
			idRange = idRange[0 : len(idRange)-1]
		}

		startStop := strings.Split(idRange, "-")
		if len(startStop) != 2 {
			panic("Invalid range format: " + idRange)
		}

		start, err := strconv.Atoi(startStop[0])
		if err != nil {
			panic(fmt.Errorf("Issue parsing start value: %w", err))
		}
		stop, err := strconv.Atoi(startStop[1])
		if err != nil {
			panic(fmt.Errorf("Issue parsing stop value: %w", err))
		}

		for id := start; id <= stop; id++ {
			idString := strconv.Itoa(id)
			idLen := len(idString)

			if idLen%2 != 0 {
				// Number must have even count of digits
				continue
			}

			if idString[0:idLen/2] == idString[idLen/2:idLen] {
				invalidTotal += id
			}
		}

		if readErr == io.EOF {
			break
		}
	}

	fmt.Println("Sum of invalid ID's:", invalidTotal)
}
