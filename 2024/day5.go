// day4.go
package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func isValidPair(order [][]int, n1, n2 int) bool {
	fmt.Println(n1, n2)
	for _, o := range order {
		if n2 == o[0] && n1 == o[1] {
			return false
		}
	}
	return true
}

func isValid(order [][]int, update []int) bool {
	for i := range len(update) {
		for j := i; j < len(update); j += 1 {
			if !isValidPair(order, update[i], update[j]) {
				return false
			}
		}
	}
	return true
}

func rearrange(order [][]int, update []int) int {
	return 0
}

func day5(input []byte) (string, string) {
	lines := bytes.Split(input, []byte{'\n'})
	re := regexp.MustCompile(`(\d{2})\|(\d{2})`)
	i := 0
	order := make([][]int, 0)
	for j, line := range lines {
		if len(line) == 0 {
			i = j
			break
		}
		matches := re.FindStringSubmatch(string(line))
		m1, _ := strconv.Atoi(matches[1])
		m2, _ := strconv.Atoi(matches[2])
		order = append(order, []int{m1, m2})
	}
	updates := make([][]int, 0)
	for i = i + 1; i < len(lines); i += 1 {
		update := make([]int, 0)
		for _, n := range strings.Split(strings.TrimSpace(string(lines[i])), ",") {
			num, _ := strconv.Atoi(n)
			update = append(update, num)
		}
		updates = append(updates, update)
	}
	p1 := 0
	p2 := 0
	for _, u := range updates {
		if isValid(order, u) {
			p1 += u[len(u)/2]
		} else {
			p2 += rearrange(order, u)
		}
	}

	return fmt.Sprintf("%d", p1), ""
}
