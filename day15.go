// day2.go
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func animate(g [][]byte, x, y int) {
	g[x][y] = '@'
	time.Sleep(500 * time.Millisecond)
	var cmd *exec.Cmd
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	for _, l := range g {
		fmt.Println(string(l))
	}
	g[x][y] = '.'
}

func day15(input []byte) (string, string) {
	parts := bytes.Split(input, []byte("\n\n"))
	gridRaw := parts[0]
	linesRaw := parts[1]
	grid := make([][]byte, 0)
	gridLines := bytes.Split(gridRaw, []byte("\n"))
	grid = append(grid, gridLines...)
	directions := (bytes.Replace(linesRaw, []byte("\n"), []byte(""), -1))
	x, y := 0, 0
	// m, n := len(grid), len(grid[0])
	// for i := range m {
	// 	for j := range n {
	// 		if grid[i][j] == '@' {
	// 			x, y = i, j
	// 		}
	// 	}
	// }
	dirMap := map[byte][2]int{
		'<': {0, -1},
		'^': {-1, 0},
		'>': {0, 1},
		'v': {1, 0},
	}

	// for _, dir := range directions {
	// 	d := dirMap[dir]
	// 	i, j := x+d[0], y+d[1]
	// 	if grid[i][j] == '.' {
	// 		x, y = i, j
	// 	} else if grid[i][j] == 'O' {
	// 		_i, _j := i, j
	// 		for grid[_i][_j] == 'O' {
	// 			_i, _j = _i+d[0], _j+d[1]
	// 		}
	// 		if grid[_i][_j] == '.' {
	// 			x, y = i, j
	// 			grid[i][j] = '.'
	// 			grid[_i][_j] = 'O'
	// 		}
	// 	}
	// }
	// p1 := 0
	// for i := range m {
	// 	for j := range n {
	// 		if grid[i][j] == 'O' {
	// 			p1 += 100*i + j
	// 		}
	// 	}
	// }

	g := make([][]byte, 0)
	for range grid {
		g = append(g, make([]byte, 0))
	}
	for i, l := range grid {
		for j, c := range l {
			if c == '#' || c == '.' {
				g[i] = append(g[i], []byte{c, c}...)
			}
			if c == 'O' {
				g[i] = append(g[i], '[')
				g[i] = append(g[i], ']')
			}
			if c == '@' {
				g[i] = append(g[i], '.')
				g[i] = append(g[i], '.')
				x, y = i, 2*j
			}
		}
	}
	g[x][y] = '.'
	for _, dir := range directions {
		d := dirMap[dir]
		i, j := x+d[0], y+d[1]
		if g[i][j] == '.' {
			x, y = i, j
			animate(g, x, y)
			continue
		}
		if g[i][j] == '#' {
			animate(g, x, y)
			continue
		}
		if (dir == '<') || (dir == '>') {
			// Search for ending position we can slide the blocks.
			_i, _j := i, j
			for g[_i][_j] == '[' || g[_i][_j] == ']' {
				_i, _j = _i+d[0], _j+d[1]
			}
			// Flip all brackets to slide them over.
			if g[_i][_j] == '.' {
				_i, _j = i, j
				for g[_i][_j] == '[' || g[_i][_j] == ']' {
					if g[_i][_j] == '[' {
						g[_i][_j] = ']'
					} else {
						g[_i][_j] = '['
					}
					_i, _j = _i+d[0], _j+d[1]
				}
				// Insert closing bracket.
				if dir == '<' {
					g[_i][_j] = '['
				} else {
					g[_i][_j] = ']'
				}
				g[i][j] = '.'
				x, y = i, j
			}
		}
		if dir == 'v' || dir == '^' {
			stack := make([]T, 0)
			stack = append(stack, T{i, j})
			stack = append(stack, otherHalf(g, T{i, j}))
			visited := make(map[T]bool)
			cache := make(map[T]byte)
			cache[T{i, j}] = '.'
			cache[otherHalf(g, T{i, j})] = '.'
			walls := false
			for len(stack) > 0 && !walls {
				u1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				u2 := otherHalf(g, u1)
				if !visited[u2] {
					stack = append(stack, u2)
					if _, ok := cache[u2]; !ok {
						cache[u2] = '.'
					}
				}
				if visited[u1] {
					continue
				} else {
					visited[u1] = true
				}
				_i, _j := u1.A+d[0], u1.B+d[1]
				if g[_i][_j] == '#' {
					walls = true
					break
				}
				if g[_i][_j] == '.' {
					cache[T{_i, _j}] = g[u1.A][u1.B]
				} else {
					cache[T{_i, _j}] = g[u1.A][u1.B]
					stack = append(stack, T{_i, _j})
				}
			}
			if !walls {
				x, y = i, j
				for loc, b := range cache {
					g[loc.A][loc.B] = b
				}
			}
		}
		animate(g, x, y)

	}
	p2 := 0
	for i, l := range g {
		for j, b := range l {
			if b == '[' {
				p2 += (100 * i) + j
			}
		}
	}
	return fmt.Sprintf("%d", p2), ""
}

func otherHalf(g [][]byte, loc T) T {
	if g[loc.A][loc.B] == ']' {
		return T{loc.A, loc.B - 1}
	} else {
		return T{loc.A, loc.B + 1}
	}
}
