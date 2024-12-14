// day2.go
package main

import (
	"bytes"
	"fmt"
)

func regionArea(grid [][]byte, visited map[T]bool, m, n, i, j int) int {
	dirs := [][]int{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}
	sides := 0
	area := 0
	ch := grid[i][j]
	stack := make([]T, 0)
	stack = append(stack, T{i, j})
	for len(stack) > 0 {

		plot := stack[len(stack)-1]
		i, j := plot.A, plot.B
		stack = stack[:len(stack)-1]
		if visited[plot] {
			continue
		} else {
			area += 1
			visited[plot] = true
			sides += countCorners(grid, m, n, i, j)
		}
		for _, d := range dirs {
			x, y := i+d[0], j+d[1]
			if inBounds(m, n, x, y) && grid[x][y] == ch {
				if !visited[T{x, y}] {
					stack = append(stack, T{x, y})
				}
			}
		}
	}
	return area * sides
}

func countCorners(grid [][]byte, m, n, i, j int) int {
	dirs := [][]int{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}
	ch := grid[i][j]
	res := 0
	// Concave corners.
	for _i := range 4 {
		_j := (_i + 1) % 4
		sides := [][]int{dirs[_i], dirs[_j]}
		cond := true
		for _, s := range sides {
			x, y := i+s[0], j+s[1]
			if inBounds(m, n, x, y) && grid[x][y] == ch {
				cond = false
			}
		}
		if cond {
			res += 1
		}
	}
	// Convex corners.
	for _i := range 4 {
		_j := (_i + 1) % 4
		d1, d2 := dirs[_i], dirs[_j]
		x1, y1 := i+d1[0], j+d1[1]
		x2, y2 := i+d2[0], j+d2[1]
		x3, y3 := i+d1[0]+d2[0], j+d1[1]+d2[1]
		if inBounds(m, n, x1, y1) && inBounds(m, n, x2, y2) && inBounds(m, n, x3, y3) {
			if grid[x1][y1] == ch && grid[x2][y2] == ch && grid[x3][y3] != ch {
				res += 1
			}
		}
	}
	return res
}

func day12(input []byte) (string, string) {
	grid := bytes.Split(input, []byte{'\n'})
	m, n := len(grid), len(grid)
	visited := make(map[T]bool)
	// True is concave, False is convex.
	p1 := 0
	for i := range m {
		for j := range n {
			if !visited[T{i, j}] {
				p1 += regionArea(grid, visited, m, n, i, j)

			}

		}
	}
	return "", fmt.Sprintf("%d", p1)
}
