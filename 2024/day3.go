// day2.go
package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func findMultiplications(input []byte) []int {
	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do(?:n't)?\(\)`)
	mulPattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	enabled := true
	// Find all matches in the input
	matches := pattern.FindAll(input, -1)

	results := make([]int, 0, len(matches))

	for _, match := range matches {
		str := string(match)
		if mulMatch := mulPattern.FindSubmatch(match); mulMatch != nil {
			if enabled {
				x, err1 := strconv.Atoi(string(mulMatch[1]))
				y, err2 := strconv.Atoi(string(mulMatch[2]))
				if err1 == nil && err2 == nil {
					results = append(results, x*y)
				}
			}
		} else if str == "do()" {
			enabled = true
		} else if str == "don't()" {
			enabled = false
		}
	}
	return results
}

func day3(input []byte) (string, string) {
	results := findMultiplications(input)
	res := 0
	for _, r := range results {
		res += r
	}
	return fmt.Sprintf("%d", res), ""
}
