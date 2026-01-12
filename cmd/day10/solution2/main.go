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

---- Part 2 ----
- Now that the machines are powered on, we need to
	worry about the joltage levels.
- Below the buttons on each machine is a big lever
	that you can use to switch the buttons from
	configuring the indicator lights to increasing
	the joltage levels.
- The indicator light diagrams are now IRRELEVANT,
	and can be ignored.
- The machines each have a set of numeric counters
	tracking its joltage levels, ONE counter per
	joltage requirement.
	- The counters are all initially set to zero.
- In this new joltage configuration mode:
	- Each button now indicates which counters it
		affects, where 0 means the first counter,
		1 means the second counter, and so on.
	- When you push a button, each listed counter is
		increased by 1.
- You can push each button as many times as you like.
- To find the answer:
	- What is the FEWEST total presses required to
		correctly configure each machine's joltage
		level counters to match the specified joltage
		requirements?
*/

/*
---- Brute Force ----
- Try combinations of button presses while increasing
	number of presses.
	- By manually incrementing num of presses, we can
		"stop low."
- Now I think everything just needs to be int slices
- Run time for test input now suspiciously high
- Suspecting there is a math solution
*/

func main() {
	var result int = 0 // Fewest presses required

	joltDiagrams := [][]int{}
	buttons := [][][]int{}

	input, err := os.Open("./cmd/day10/input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		joltString := parts[len(parts)-1]
		joltString = joltString[1 : len(joltString)-1] // Cut off curly brackets

		joltNums := strings.Split(joltString, ",")

		parsedDiagram := []int{}
		for _, str := range joltNums {
			num, err := strconv.Atoi(str)
			if err != nil {
				panic("Failed to parse joltage value: " + err.Error())
			}
			parsedDiagram = append(parsedDiagram, num)
		}
		joltDiagrams = append(joltDiagrams, parsedDiagram)

		buttons = append(buttons, [][]int{})

		for i := 1; i < len(parts)-1; i++ {
			buttonStr := parts[i]
			buttonStr = buttonStr[1 : len(buttonStr)-1] // cut off parentheses
			parsedButton := []int{}

			nums := strings.Split(buttonStr, ",")

			for _, num := range nums {
				joltPos, err := strconv.Atoi(num)
				if err != nil {
					panic(fmt.Errorf("Failed to parse button's wiring: %w", err))
				}
				parsedButton = append(parsedButton, joltPos)
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
		fmt.Println("Machine:", machine+1, "/ 170")
		combinations := [][]int{}

		for _, button := range row {
			// Depth 1

			joltCounts := []int{}
			for range len(joltDiagrams[machine]) {
				joltCounts = append(joltCounts, 0)
			}

			for _, joltPos := range button {
				joltCounts[joltPos] += 1
			}

			if slices.Equal(joltCounts, joltDiagrams[machine]) {
				result += 1
				continue MachineLoop
			}

			combinations = append(combinations, joltCounts)
		}

		for depth := 2; true; depth++ {
			fmt.Println("Depth:", depth)
			fmt.Println("Combinations:", len(combinations))
			// Combinations become unwieldly around depth = 10
			// Find a way to use math instead
			newCombinations := [][]int{}

			for _, button := range row {
				for _, combo := range combinations {
					joltCounts := make([]int, len(combo))
					copy(joltCounts, combo)

					for _, joltPos := range button {
						joltCounts[joltPos] += 1
					}

					if slices.Equal(joltCounts, joltDiagrams[machine]) {
						result += depth
						continue MachineLoop
					}

					newCombinations = append(
						newCombinations, joltCounts,
					)
				}
			}

			combinations = newCombinations
		}
	}

	fmt.Println("Fewest possible button presses to get the correct joltage levels:", result)
}
