// day4.go
package main

import (
	"bytes"
	"fmt"
)

func inBounds(m, n, i, j int) bool {
	return 0 <= i && i < m && 0 <= j && j < n
}

func day6(input []byte) (string, string) {
	lines := bytes.Split(input, []byte{'\n'})
	x, y := 0, 0
	m, n := len(lines), len(lines[0])
	for i := range m {
		for j := range n {
			if lines[i][j] == '^' {
				x, y = i, j
				// goto considered useful.
				goto exit_start_pos
			}
		}
	}
exit_start_pos:
	lines[x][y] = 'X'
	dirs := [][]int{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}
	d := 0
	res := 1
	p2 := 0
	fmt.Println(x, y)
	fmt.Println(m, n)
	for inBounds(m, n, x, y) {
		fmt.Println(x, y)
		i, j := x+dirs[d][0], y+dirs[d][1]
		if inBounds(m, n, i, j) {
			if lines[i][j] == '.' {
				res += 1
				lines[i][j] = 'X'
				x, y = i, j
			} else if lines[i][j] == 'X' {
				x, y = i, j
				p2 += 1
			} else {
				d = (d + 1) % 4
			}
		} else {
			break
		}

	}
	return fmt.Sprintf("%d", res), fmt.Sprintf("%d", p2)
}
