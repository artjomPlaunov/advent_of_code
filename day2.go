// day2.go
package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

func removeElement(r []int, i int) []int {
	res := make([]int, 0, cap(r))
	res = append(res, r[:i]...)
	res = append(res, r[i+1:]...)
	return res
}

func isSafe(r []int) (bool, int, int) {
	if r[0] == r[1] {
		return false, 0, 1
	}
	inc := true
	if r[0] > r[1] {
		inc = false
	}
	for i := 1; i < len(r); i++ {
		if r[i] == r[i-1] || inc && r[i-1] > r[i] || !inc && r[i-1] < r[i] || math.Abs(float64(r[i])-float64(r[i-1])) > 3 {
			return false, i - 1, i
		}
	}
	return true, -1, -1
}

func day2(input []byte) (string, string) {
	lines := bytes.Split(input, []byte{'\n'})
	var reports [][]int
	for _, line := range lines {
		report := make([]int, 0)
		ns := bytes.Split(line, []byte{' '})
		for _, n := range ns {
			n, _ := strconv.Atoi(string(n))
			report = append(report, n)
		}
		reports = append(reports, report)
	}
	validCount := 0
	for _, r := range reports {
		safe1, i1, i2 := isSafe(r)
		rem := []int{i1, i2}
		if safe1 {
			validCount++
		} else {

			for _, idx := range rem {
				r1 := removeElement(r, idx)
				s1, _, _ := isSafe(r1)
				if s1 {
					validCount++
					break
				}
			}
		}
	}
	return fmt.Sprintf("%d", validCount), ""
}
