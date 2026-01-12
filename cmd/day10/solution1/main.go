package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
- Elves in one of the toy factories can't figure
	out the initialization procedure for the machines.
- The manual (your input) describes one machine per
	line:
	- Each line contains:
		- A single indicator light diagram in
			[square brackets]
		- One or more button wiring schematics in
			(parentheses)
			- Each parentheses pair represents a SINGLE
				button, that affects MULTIPLE lights.
		- Joltage requirements in {curly braces}.
- Because none of the machines are running, the
	joltage requirements are IRRELEVANT and can be
	safely ignored.
- To start a machine:
	- Its indicator lights must match those shown in
		the diagram, where . means off and # means on.
- The machine has the number of indicator lights
	shown, BUT:
	- Its indicator lights are all INITIALLY
	OFF.
- You can TOGGLE the state of indicator lights by
	pushing any of the listed buttons:
	- Each button lists which indicator lights it
		toggles, where 0 means the first light,
		1 means the second light, and so on.
	- When you push a button, each listed indicator
		light either turns on (if it was off) or turns
		off (if it was on).
	- You have to push each button an integer number
		of times; there's no such thing as "0.5 presses"
		(nor can you push a button a negative number of
		times).
- You can push each button as many times as you like.
- To find the answer:
	- What is the FEWEST number of button presses
		required to correctly configure the indicator
		lights on ALL of the machines?
*/

// N = 170 machines
/*
---- Brute Force ----
- Try combinations of button presses while increasing
	number of presses.
	- By manually incrementing num of presses, we can
		"stop low."
- Binary representation of lights should make testing
	combinations much easier
	- As easy to compare to diagram as string would be
	- "Combining" buttons is as simple as bitwise XOR
*/

func main() {
	var result int = 0 // Fewest presses required

	// These ints are bit representations of lights
	lightDiagrams := []int{}
	buttons := [][]int{}

	input, err := os.Open("./cmd/day10/input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		lightDiagram := parts[0][1 : len(parts[0])-1] // Cut off square brackets
		lenDiagram := len(lightDiagram)

		parsedDiagram := 0
		for i, char := range lightDiagram {
			if char == '#' {
				parsedDiagram |= 1 << (lenDiagram - i - 1)
			}
		}
		lightDiagrams = append(lightDiagrams, parsedDiagram)

		buttons = append(buttons, []int{})

		for i := 1; i < len(parts)-1; i++ {
			button := parts[i]
			parsedButton := 0

			nums := strings.Split(button[1:len(button)-1], ",")

			for _, num := range nums {
				lightPos, err := strconv.Atoi(num)
				if err != nil {
					panic(fmt.Errorf("Failed to parse button's wiring: %w", err))
				}
				parsedButton |= 1 << (lenDiagram - lightPos - 1)
			}

			buttons[len(buttons)-1] = append(
				buttons[len(buttons)-1], parsedButton,
			)
		}
	}

	// Try depth = 1 all combinations
	/*
		--- Account for Higher Depth --
		- Either by recursion or memoization
			- Maybe both
		- Def need memoization because there is no way
			to check a higher depth in the MIDDLE of
			checking a lower one.
			- So create "combinations" array
		- Recursion may end up being unnecessary then
			- Yup, memoization did the trick.
	*/
MachineLoop:
	for machine, row := range buttons {
		for _, button := range row {
			// Lights start all off
			// So we can take button to be initial state
			if button == lightDiagrams[machine] {
				result += 1
				continue MachineLoop
			}
		}

		combinations := make([]int, len(row))
		copy(combinations, row)

		for depth := 2; true; depth++ {
			newCombinations := []int{}

			for _, button := range row {
				for _, combo := range combinations {
					if button == combo {
						continue
					}

					if button^combo == lightDiagrams[machine] {
						result += depth
						continue MachineLoop
					}

					newCombinations = append(
						newCombinations, button^combo,
					)
				}
			}

			combinations = newCombinations
		}
	}

	// 465 -> Too High
	// 457 -> CORRECT
	fmt.Println("Fewest possible button presses to turn on all of the machines:", result)
}
