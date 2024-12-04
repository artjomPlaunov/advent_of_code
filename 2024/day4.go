// day4.go
package main

import (
	"bytes"
	"fmt"
)

func matches(i, j int, grid [][]byte, word []byte) int {
	dirs := [][]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}
	res := 0
	for _, d := range dirs {
		found := true
		for l := range len(word) {
			x, y := (d[0]*l)+i, (d[1]*l)+j
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
				if word[l] != grid[x][y] {
					found = false
					break
				}
			} else {
				found = false
				break
			}
		}
		if found {
			res += 1
		}
	}
	return res
}

func day4(input []byte) (string, string) {
	lines := bytes.Split(input, []byte{'\n'})
	res := 0
	n, m := len(lines), len(lines[0])
	for i := range n {
		for j := range m {
			res += matches(i, j, lines, []byte("XMAS"))
		}
	}
	return fmt.Sprintf("%d", res), ""
}
