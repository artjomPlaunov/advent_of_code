// day2.go
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func dp(patterns map[string]bool, design string, idx int, cache map[int]int) {
	_, ok := cache[idx]
	if ok {
		return
	}
	for i := idx; i < len(design); i++ {
		if patterns[design[idx:i+1]] {

			if i+1 == len(design) {
				cache[idx] += 1
				return
			} else {
				dp(patterns, design, i+1, cache)
				if cache[i+1] > 0 {
					cache[idx] += cache[i+1]
				}
			}
		}
	}
}

func day19(input []byte) (string, string) {
	lines := bytes.Split(input, []byte("\n\n"))
	designs := bytes.Split(lines[1], []byte{'\n'})
	values := strings.Split(string(lines[0]), ", ")
	patterns := make(map[string]bool)
	for _, v := range values {
		patterns[v] = true
	}
	p1 := 0
	for _, d := range designs {
		cache := make(map[int]int)
		dp(patterns, string(d), 0, cache)
		p1 += cache[0]
	}

	return fmt.Sprintf("%d", p1), ""
}
