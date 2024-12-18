// day2.go
package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type N struct {
	X, Y int
	Dist int
}

type NH []N

func (h NH) Len() int { return len(h) }

func (h NH) Less(i, j int) bool {
	return h[i].Dist < h[j].Dist
}

func (h NH) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NH) Push(x interface{}) {
	*h = append(*h, x.(N))
}
func (h *NH) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func dijkstra(bytes map[T]bool) int {
	dirs := [][]int{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}
	n := 71
	maxDist := math.MaxInt
	minDists := make(map[T]int)
	h := &NH{}
	heap.Init(h)
	for i := range n {
		for j := range n {
			if i != 0 || j != 0 {
				minDists[T{i, j}] = maxDist
				heap.Push(h, N{0, 0, maxDist})
			}
		}
	}

	heap.Push(h, N{0, 0, 0})
	minDists[T{0, 0}] = 0
	res := maxDist
	// for i := range n {
	// 	for j := range n {
	// 		if _, ok := bytes[T{i, j}]; ok {
	// 			fmt.Print(string('#'))
	// 		} else {
	// 			fmt.Print(string('.'))
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	for h.Len() > 0 {
		u := heap.Pop(h).(N)
		if u.X == n-1 && u.Y == n-1 {
			res = minDists[T{n - 1, n - 1}]
			return res
		}
		if minDists[T{u.X, u.Y}] < u.Dist {
			continue
		}
		for _, d := range dirs {

			x, y := u.X+d[0], u.Y+d[1]
			if bytes[T{x, y}] {
				continue
			}
			if u.Dist+1 < minDists[T{x, y}] {
				minDists[T{x, y}] = u.Dist + 1
				heap.Push(h, N{x, y, u.Dist + 1})
			}
		}
	}
	return res
}

func day18(input []byte) (string, string) {

	lines := bytes.Split(input, []byte{'\n'})
	p2 := T{0, 0}
	bytes := make(map[T]bool)
	for _, l := range lines {

		nums := strings.Split(string(l), ",")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		bytes[T{n2, n1}] = true
		if dijkstra(bytes) == math.MaxInt {
			p2 = T{n1, n2}
			break
		}
		//fmt.Println(n2, n1)
	}
	bytes = make(map[T]bool)
	p1 := dijkstra(bytes)

	return fmt.Sprintf("%d", p1), fmt.Sprintf("%d,%d", p2.A, p2.B)
}
