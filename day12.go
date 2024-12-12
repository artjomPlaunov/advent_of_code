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
		}
		for _, d := range dirs {
			x, y := i+d[0], j+d[1]
			if inBounds(m, n, x, y) && grid[x][y] == ch {
				if !visited[T{x, y}] {
					stack = append(stack, T{x, y})
				}
			} //else {
			//	perim += 1
			//}
		}
	}
	//return perim * area
	return area
}

func isCorner(grid [][]byte, m, n, i, j int) bool {
	dirs := [][]int{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}
	ch := grid[i][j]
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
			return true
		}
	}

	return false
}

func day12(input []byte) (string, string) {
	grid := bytes.Split(input, []byte{'\n'})
	m, n := len(grid), len(grid)
	visited := make(map[T]bool)
	sidesVisited := make(map[T]bool)
	p1 := 0
	for i := range m {
		for j := range n {
			//area, side := 0, 0
			if !visited[T{i, j}] {
				// Use part1 to get area, sides will be calculated separately.
				p1 += regionArea(grid, visited, m, n, i, j)
			}
			if !sidesVisited[T{i, j}] && isCorner(grid, m, n, i, j) {
				fmt.Println(i, j)
			}

		}
	}
	return "", ""
}
