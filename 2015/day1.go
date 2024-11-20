// day1.go
package main

import "strconv"

func day1(input []byte) (string, string) {
	floor := 0
	firstBasement := 0
	foundBasement := false

	// Process each character
	for i, b := range input {
		if b == '(' {
			floor++
		} else {
			floor--
		}

		// Check for first basement entry (part 2)
		if floor < 0 && !foundBasement {
			firstBasement = i + 1
			foundBasement = true
		}
	}

	// Convert results to strings and return
	return strconv.Itoa(floor), strconv.Itoa(firstBasement)
}
