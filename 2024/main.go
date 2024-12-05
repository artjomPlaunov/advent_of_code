// main.go
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// DayFunc is the function signature that all daily solutions must follow
type DayFunc func([]byte) (string, string)

// solutions maps day numbers to their solving functions
var solutions = map[int]DayFunc{
	1: day1,
	2: day2,
	3: day3,
	4: day4,
	5: day5,
}

func main() {
	// Check if a day number was provided
	if len(os.Args) != 2 {
		log.Fatal("Please provide day number as argument")
	}

	// Convert the day argument to a number
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Invalid day number:", err)
	}

	// Get the solving function for this day
	solution, exists := solutions[day]
	if !exists {
		log.Fatal("Solution not implemented for day", day)
	}

	// Read input file from the input directory
	input, err := os.ReadFile(fmt.Sprintf("input/day%d.txt", day))
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	// Run the solution and print results
	part1, part2 := solution(input)
	fmt.Printf("Day %d:\n", day)
	fmt.Printf("Part 1: %s\n", part1)
	fmt.Printf("Part 2: %s\n", part2)
}
