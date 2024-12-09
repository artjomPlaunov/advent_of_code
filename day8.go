// day2.go
package main

import (
	"bytes"
	"fmt"
)

func numAntinodes(a, b T, m, n int, seen map[T]bool) int {
	dx := a.A - b.A
	dy := a.B - b.B
	res := 0
	p1, p2 := T{a.A + dx, a.B + dy}, T{b.A - dx, b.B - dy}
	if inBounds(m, n, p1.A, p1.B) {
		if !seen[p1] {
			seen[p1] = true
			res += 1
		}
	}
	if inBounds(m, n, b.A-dx, b.B-dy) {
		if !seen[p2] {
			seen[p2] = true
			res += 1
		}
	}
	return res
}
func numAntinodes2(a, b T, m, n int, seen map[T]bool) int {
	dx := a.A - b.A
	dy := a.B - b.B
	res := 0
	p1 := T{a.A + dx, a.B + dy}
	fmt.Println("here", a, b)
	for inBounds(m, n, p1.A, p1.B) {
		fmt.Println(p1)
		if !seen[p1] {
			seen[p1] = true
			res += 1
		}
		p1.A += dx
		p1.B += dy
	}
	p1 = T{b.A - dx, b.B - dy}
	for inBounds(m, n, p1.A, p1.B) {
		fmt.Println(p1)
		if !seen[p1] {
			seen[p1] = true
			res += 1
		}
		p1.A -= dx
		p1.B -= dy
	}
	return res
}

func day8(input []byte) (string, string) {
	lines := bytes.Split(input, []byte{'\n'})
	m, n := len(lines), len(lines[0])
	nodes := make(map[byte][]T)
	for i, l := range lines {
		for j, b := range l {
			if b != '.' {
				if _, ok := nodes[b]; ok {
					nodes[b] = append(nodes[b], T{i, j})
				} else {
					nodes[b] = []T{{i, j}}
				}
			}
		}
	}
	antinodes := 0
	antinodes2 := 0
	seen2 := make(map[T]bool)
	seen := make(map[T]bool)
	for _, positions := range nodes {
		for i := range len(positions) {
			if len(positions) > 2 && !seen2[positions[i]] {
				seen2[positions[i]] = true
				antinodes2 += 1
			}
			for j := i + 1; j < len(positions); j++ {
				antinodes += numAntinodes(positions[i], positions[j], m, n, seen)
				antinodes2 += numAntinodes2(positions[i], positions[j], m, n, seen2)
			}
		}
	}
	return fmt.Sprintf("%d", antinodes), fmt.Sprintf("%d", antinodes2)
}
