// day2.go
package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func mix(n, m int) int {
	return n ^ m
}

func prune(n int) int {
	return n % 16777216
}

func step(n int) int {
	tmp := n * 64
	n = mix(n, tmp)
	n = prune(n)
	tmp = n / 32
	n = mix(n, tmp)
	n = prune(n)
	tmp = n * 2048
	n = mix(n, tmp)
	return prune(n)
}

func day22(input []byte) (string, string) {
	lines := bytes.Split(input, []byte{'\n'})
	nums := make([]int, 0)
	for _, l := range lines {
		n, _ := strconv.Atoi(string(l))
		nums = append(nums, n)
	}
	p1 := 0

	changes := make([][]T, 0)
	for _, n := range nums {
		cs := make([]T, 0)
		for range 2000 {
			tmp := n % 10
			n = step(n)
			cs = append(cs, T{(n % 10) - tmp, n % 10})
		}
		changes = append(changes, cs)
		p1 += n
	}
	cache := make(map[int]map[T4]int)
	for j, c := range changes {
		cache[j] = make(map[T4]int)
		for i := 0; i < len(c)-3; i++ {
			val := c[i+3].B
			deltas := T4{c[i].A, c[i+1].A, c[i+2].A, c[i+3].A}
			if _, ok := cache[j][deltas]; !ok {
				cache[j][deltas] = val
			}
		}
	}
	p2 := 0
	for a := -9; a <= 9; a++ {
		for b := -9; b <= 9; b++ {
			for C := -9; C <= 9; C++ {
				for d := -9; d <= 9; d++ {
					res := 0
					for _, c := range cache {
						res += c[T4{a, b, C, d}]
					}
					p2 = max(p2, res)
				}
			}
		}
	}
	return fmt.Sprintf("%d", p1), fmt.Sprintf("%d", p2)
}
