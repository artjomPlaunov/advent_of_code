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

	minDists := make(map[T]map[int]int)
	for i := range m {
		for j := range n {
			minDists[T{i, j}] = make(map[int]int)
			for k := range 4 {
				minDists[T{i, j}][k] = 99999999999999
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
	p1 := 0
	for h.Len() > 0 {
		u := heap.Pop(h).(Node)
		if u.A == e.A && u.B == e.B {
			p1 = minDists[T{u.A, u.B}][u.Dir]
			break
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

	return fmt.Sprintf("%d", p1), ""
}
