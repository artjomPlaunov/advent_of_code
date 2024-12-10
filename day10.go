package main

import (
	"bytes"
	"fmt"
)

func neighbors(i, j, m, n int) []T {
	dirs := []T{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	res := make([]T, 0)
	for _, d := range dirs {
		x, y := i+d.A, j+d.B
		if inBounds(m, n, x, y) {
			res = append(res, T{x, y})
		}
	}
	return res
}

// I initially forgot to have a visited set for part1; this was necessary because we
// just wanted the number of peaks reachable from the trailhead, so we didn't care about
// various paths to get there. Once we reach a node and explored its neighbors, that node
// is marked as "visited" and we are done with it.

// By a funny coincidence, I forgot to have it and realized that what I was doing was actually
// counting all the possible paths to the summit, i.e., part2 of the puzzle.
func countTrailHeads(grid [][]byte, i, j, m, n int) int {
	stk := make([]T, 0)
	stk = append(stk, T{i, j})
	//visited := make(map[T]bool)
	res := 0
	for len(stk) > 0 {
		fmt.Println(stk)
		l := len(stk)
		i, j := stk[l-1].A, stk[l-1].B
		stk = stk[:l-1]
		//visited[T{i, j}] = true
		if int(grid[i][j]-'0') == 9 {
			res += 1
			continue
		}

		for _, n := range neighbors(i, j, m, n) {
			x, y := n.A, n.B
			//if int(grid[x][y]-'0')-1 == int(grid[i][j]-'0') && !visited[T{x, y}] {
			if int(grid[x][y]-'0')-1 == int(grid[i][j]-'0') {
				stk = append(stk, T{x, y})
			}
		}
	}
	return res
}

func day10(input []byte) (string, string) {
	grid := bytes.Split(input, []byte{'\n'})
	p1 := 0
	m, n := len(grid), len(grid[0])
	for i := range m {
		for j := range n {
			if int(grid[i][j]-'0') == 0 {
				p1 += countTrailHeads(grid, i, j, m, n)
				fmt.Println(i, j, p1)
			}
		}
	}

	return fmt.Sprintf("%d", p1), ""
}
