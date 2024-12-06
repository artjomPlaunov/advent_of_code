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
	startx, starty := 0, 0
	m, n := len(lines), len(lines[0])
	for i := range m {
		for j := range n {
			if lines[i][j] == '^' {
				startx, starty = i, j
				// goto considered useful.
				goto exit_start_pos
			}
		}
	}
exit_start_pos:

	lines[startx][starty] = 'X'
	dirs := [][]int{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}
	d := 0
	p2 := 0

	for row := range m {
		for col := range n {
			if (row != startx || col != starty) && lines[row][col] == '.' {
				grid := deepCopy(lines)
				grid[row][col] = '#'
				x, y := startx, starty
				d = 0
				visited := make(map[T]int)
				for inBounds(m, n, x, y) {
					if visited[T{x, y}] > 4 {
						p2 += 1
						break
					}
					i, j := x+dirs[d][0], y+dirs[d][1]
					if inBounds(m, n, i, j) {
						if grid[i][j] == '.' {
							grid[i][j] = 'X'
							visited[T{i, j}] += 1
							x, y = i, j
						} else if grid[i][j] == 'X' {
							visited[T{i, j}] += 1
							x, y = i, j
						} else {
							d = (d + 1) % 4
						}
					} else {
						break
					}

				}
			}
		}
	}
	x, y := startx, starty
	d = 0
	p1 := 1
	for inBounds(m, n, x, y) {
		i, j := x+dirs[d][0], y+dirs[d][1]
		if inBounds(m, n, i, j) {
			if lines[i][j] == '.' {
				p1 += 1
				lines[i][j] = 'X'
				x, y = i, j
			} else if lines[i][j] == 'X' {
				x, y = i, j
			} else {
				d = (d + 1) % 4
			}
		} else {
			break
		}

	}
	return fmt.Sprintf("%d", p1), fmt.Sprintf("%d", p2)
}
