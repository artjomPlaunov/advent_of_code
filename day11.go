package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func evenDigits(n int) bool {
	if n == 0 {
		return false
	}
	i := 0
	for n > 0 {
		n = n / 10
		i += 1
	}
	if i%2 == 0 {
		return true
	}
	return false
}

func split(n int) []int {
	numStr := strconv.Itoa(n)
	m := len(numStr) / 2
	lhs := numStr[:m]
	rhs := numStr[m:]
	n1, _ := strconv.Atoi(strings.TrimLeft(lhs, "0"))
	n2, _ := strconv.Atoi(strings.TrimLeft(rhs, "0"))
	return []int{n1, n2}
}

func applyRule(n int) []int {
	//res := make([]int, 0)
	if n == 0 {
		return []int{1}
	}
	isEven := evenDigits(n)
	if isEven {
		return split(n)
	} else {
		return []int{2024 * n}
	}
}

func day11(input []byte) (string, string) {
	nums := bytes.Split(input, []byte{' '})
	stones := make(map[int]int)
	for _, n := range nums {

		num, _ := strconv.Atoi(string(n))
		stones[num] = 1
	}
	for range 75 {
		aux := make(map[int]int)
		for n, count := range stones {
			for _, d := range applyRule(n) {
				aux[d] += count
			}
		}
		for key := range stones {
			delete(stones, key)
		}
		for key, value := range aux {
			stones[key] = value
		}
	}
	p1 := 0
	for _, value := range stones {
		p1 += value
	}

	return fmt.Sprintf("%d", p1), ""
}
