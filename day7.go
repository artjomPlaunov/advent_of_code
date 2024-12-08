// day7.go
package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

var numRegex = regexp.MustCompile(`\d+`)

func isValidEq(eq []int) bool {
	n := eq[0]
	numOps := len(eq) - 2
	permutations := (1 << numOps)
	for p := range permutations {
		s := eq[1]
		for i := 2; i < len(eq); i++ {
			if p&1 == 1 {
				s *= eq[i]
			} else {
				s += eq[i]
			}
			p >>= 1
		}
		if s == n {
			return true
		}
	}
	return false
}

func concatenate(n, x int) int {
	res, _ := strconv.Atoi(strconv.Itoa(n) + strconv.Itoa(x))
	return res
}

func isValidEq2(eq []int) bool {
	n := eq[0]
	numOps := len(eq) - 2
	permutations := generatePermutations(0, numOps)
	for _, p := range permutations {
		s := eq[1]
		for i := 2; i < len(eq); i++ {
			op := p[i-2]
			if op == 0 {
				s += eq[i]
			} else if op == 1 {
				s *= eq[i]
			} else if op == 2 {
				s = concatenate(s, eq[i])
			}
		}
		if s == n {
			return true
		}
	}
	return false
}

func generatePermutations(i, n int) [][]int {
	if i == n-1 {
		return [][]int{{0}, {1}, {2}}
	} else {
		ps := generatePermutations(i+1, n)
		res := make([][]int, 0)
		for _, p := range ps {
			for i := range 3 {
				perm := make([]int, len(p))
				copy(perm, p)
				perm = append(perm, i)
				res = append(res, perm)
			}
		}
		return res
	}
}

func day7(input []byte) (string, string) {
	lines := bytes.Split(input, []byte{'\n'})
	equations := make([][]int, len(lines))

	for i, l := range lines {
		matches := numRegex.FindAll(l, -1)
		nums := make([]int, len(matches))
		for j, m := range matches {
			nums[j], _ = strconv.Atoi(string(m))
		}
		equations[i] = nums
	}
	p1 := 0
	p2 := 0
	for _, e := range equations {
		if isValidEq(e) {
			p1 += e[0]
		}
		if isValidEq2(e) {
			p2 += e[0]
		}
	}

	return fmt.Sprintf("%d", p1), fmt.Sprintf("%d", p2)
}
