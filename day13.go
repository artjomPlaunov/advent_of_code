// day2.go
package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

func day13(input []byte) (string, string) {
	grid := bytes.Split(input, []byte{'\n'})
	machines := make([][]int, 0)
	pattern := regexp.MustCompile(`(\d+).*?(\d+)`)
	for _, l := range grid {
		matches := pattern.FindStringSubmatch(string(l))
		if len(matches) == 3 {
			m1, _ := strconv.Atoi(matches[1])
			m2, _ := strconv.Atoi(matches[2])
			machines = append(machines, []int{m1, m2})
		}
	}
	p1 := 0
	for i := 0; i < len(machines); i += 3 {
		a := machines[i]
		b := machines[i+1]
		goal := machines[i+2]
		res := 401
		for i := range 200 {
			for j := range 200 {
				if (a[0]*i)+(b[0]*(j)) == goal[0] {
					if (a[1]*i)+(b[1]*j) == goal[1] {
						res = min(res, 3*i+j)
					}
				}
			}
		}

		if res < 401 {
			p1 += res
			fmt.Println(i, res)
		}
	}
	p2 := 0
	for i := 0; i < len(machines); i += 3 {
		a1 := machines[i][0]
		b1 := machines[i+1][0]
		a2 := machines[i][1]
		b2 := machines[i+1][1]
		c1 := machines[i+2][0] + 10000000000000
		c2 := machines[i+2][1] + 10000000000000
		a, b := 0, 0
		n := c1*b2 - b1*c2
		m := a1*b2 - b1*a2
		if n%m == 0 {
			a = n / m
		} else {
			continue
		}
		n = a1*c2 - c1*a2
		m = a1*b2 - b1*a2
		if n%m == 0 {
			b = n / m
		} else {
			continue
		}
		p2 += (3*a + b)
	}

	return fmt.Sprintf("%d", p1), fmt.Sprintf("%d", p2)
}
