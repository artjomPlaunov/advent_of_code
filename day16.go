// day2.go
package main

import (
	"bytes"
	"container/heap"
	"fmt"
)

type Node struct {
	A, B int
	Dir  int
	Dist int
}

type NodeHeap []Node

func (h NodeHeap) Len() int { return len(h) }

func (h NodeHeap) Less(i, j int) bool {
	return h[i].Dist < h[j].Dist
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}
func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func day16(input []byte) (string, string) {
	dirs := [][]int{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}
	lines := bytes.Split(input, []byte{'\n'})
	m, n := len(lines), len(lines[0])
	maxDist := 99999999999999
	minDists := make(map[T]map[int]int)
	for i := range m {
		for j := range n {
			minDists[T{i, j}] = make(map[int]int)
			for k := range 4 {
				minDists[T{i, j}][k] = maxDist
			}
		}
	}
	s, e := T{m - 2, 1}, T{1, n - 2}
	lines[m-2][1] = '.'
	lines[1][n-2] = '.'
	h := &NodeHeap{}
	heap.Init(h)
	heap.Push(h, Node{s.A, s.B, 1, 0})
	minDists[T{s.A, s.B}][1] = 0
	p1 := maxDist
	for h.Len() > 0 {
		u := heap.Pop(h).(Node)
		if u.A == e.A && u.B == e.B {
			if minDists[T{u.A, u.B}][u.Dir] < p1 {
				p1 = minDists[T{u.A, u.B}][u.Dir]
			}
			continue
		}
		if minDists[T{u.A, u.B}][u.Dir] < u.Dist {
			continue
		}
		for _, i := range []int{(u.Dir - 1) % 4, (u.Dir + 1) % 4, u.Dir} {
			if i < 0 {
				i = 3
			}
			d := dirs[i]
			x, y := u.A+d[0], u.B+d[1]
			cost := 1000
			if i == u.Dir {
				cost = 1
				if lines[x][y] == '.' {
					if u.Dist+cost < minDists[T{x, y}][i] {
						minDists[T{x, y}][i] = u.Dist + cost
						heap.Push(h, Node{x, y, i, u.Dist + cost})
					}
				}
			} else {

				if u.Dist+cost < minDists[T{u.A, u.B}][i] {
					minDists[T{u.A, u.B}][i] = u.Dist + cost
					heap.Push(h, Node{u.A, u.B, i, u.Dist + cost})
				}
			}
		}
	}
	visited := make(map[T]map[int]bool)
	marked := make(map[T]bool)
	for i := range m {
		for j := range n {
			visited[T{i, j}] = make(map[int]bool)
		}
	}
	stack := make([][]int, 0)
	stack = append(stack, []int{e.A, e.B, 0, 0})
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		x, y, d, dist := p[0], p[1], p[2], p[3]
		stack = stack[:len(stack)-1]
		visited[T{x, y}][d] = true
		minDist := minDists[T{x, y}][d]
		if minDist == maxDist {
			continue
		}
		if minDist+dist == p1 {
			marked[T{x, y}] = true
		} else {
			continue
		}
		for _, i := range []int{(d + 1) % 4, (d - 1) % 4, d} {
			if i < 0 {
				i = 3
			}
			if i == d {
				x, y = x-dirs[d][0], y-dirs[d][1]
				if !visited[T{x, y}][d] {
					stack = append(stack, []int{x, y, d, dist + 1})
				}
			} else {
				if !visited[T{x, y}][i] {
					stack = append(stack, []int{x, y, i, dist + 1000})
				}
			}
		}
	}
	p2 := 0
	for range marked {
		p2 += 1
	}

	fmt.Println(minDists[T{e.A, e.B}])
	return fmt.Sprintf("%d", p1), fmt.Sprintf("%d", p2)
}
