// day2.go
package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

func day14(input []byte) (string, string) {
	lines := bytes.Split(input, []byte{'\n'})
	robots := make([][]int, 0)

	re := regexp.MustCompile(`p=([-]?\d+),([-]?\d+)\s*v=([-]?\d+),([-]?\d+)`)
	for _, l := range lines {
		m := re.FindStringSubmatch(string(l))
		if len(m) == 5 {
			m1, _ := strconv.Atoi(m[1])
			m2, _ := strconv.Atoi(m[2])
			m3, _ := strconv.Atoi(m[3])
			m4, _ := strconv.Atoi(m[4])
			robots = append(robots, []int{m1, m2, m3, m4})
		}
	}
	mn := 9999999999999
	m, n := 103, 101
	second := 0
	// for idx, r := range robots {
	// 	i, j := r[1], r[0]
	// 	dx, dy := r[3], r[2]
	// 	x := (((i + 100*dx) % m) + m) % m
	// 	y := (((j + 100*dy) % n) + n) % n
	// 	robots[idx][0] = y
	// 	robots[idx][1] = x
	// }
	// q1, q2, q3, q4 := 0, 0, 0, 0
	// for _, r := range robots {

	// 	if (r[0] < n/2) && (r[1] < m/2) {
	// 		q1 += 1
	// 	}
	// 	if (r[0] > (n / 2)) && (r[1] < m/2) {
	// 		q2 += 1
	// 	}
	// 	if (r[0] < n/2) && (r[1] > (m / 2)) {
	// 		q3 += 1
	// 	}
	// 	if (r[0] > (n / 2)) && (r[1] > (m / 2)) {
	// 		q4 += 1
	// 	}
	// }
	for {
		second += 1
		for idx, r := range robots {
			i, j := r[1], r[0]
			dx, dy := r[3], r[2]
			x := (((i + dx) % m) + m) % m
			y := (((j + dy) % n) + n) % n
			robots[idx][0] = y
			robots[idx][1] = x
		}

		q1, q2, q3, q4 := 0, 0, 0, 0
		for _, r := range robots {

			if (r[0] < n/2) && (r[1] < m/2) {
				q1 += 1
			}
			if (r[0] > (n / 2)) && (r[1] < m/2) {
				q2 += 1
			}
			if (r[0] < n/2-10) && (r[1] > (m / 2)) {
				q3 += 1
			}
			if (r[0] > (n / 2)) && (r[1] > (m / 2)) {
				q4 += 1
			}
		}
		fact := q1 * q2 * q3 * q4

		if fact < mn {
			drawing := make([][]string, 103)
			for i := range drawing {
				drawing[i] = make([]string, 101)
				for j := range drawing[i] {
					drawing[i][j] = "."
				}
			}
			for _, r := range robots {
				drawing[r[1]][r[0]] = "1"
			}
			mn = fact
			for _, l := range drawing {
				fmt.Println(l)
			}
			fmt.Println(second)
		}
	}

	return "", ""
}
