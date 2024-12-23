// day23.go
package main

import (
	"bytes"
	"fmt"
	"slices"
	"sort"
	"strings"
)

func day23(input []byte) (string, string) {
	lines := bytes.Split(input, []byte{'\n'})
	graph := make(map[string][]string)
	for _, l := range lines {
		nodes := bytes.Split(l, []byte{'-'})
		u, v := string(nodes[0]), string(nodes[1])
		if _, ok := graph[u]; !ok {
			graph[u] = make([]string, 0)
		}
		graph[u] = append(graph[u], v)
		if _, ok := graph[v]; !ok {
			graph[v] = make([]string, 0)
		}
		graph[v] = append(graph[v], u)
	}
	p1 := 0
	seen := make(map[string]bool)
	for u, vs := range graph {
		if u[0] == 't' {
			for i, v1 := range vs {
				for j, v2 := range vs {
					if i != j && slices.Contains(graph[v1], v2) {
						ss := []string{u, v1, v2}
						sort.Strings(ss)
						s := strings.Join(ss, ",")
						if _, ok := seen[s]; !ok {
							seen[s] = true
							p1++
						}
					}
				}
			}
		}
	}
	components := make(map[string][]string)
	for u := range graph {
		components[u] = []string{u}
	}
	for u := range graph {
		for v := range graph {
			if !slices.Contains(components[u], v) {
				connected := true
				for _, s := range components[u] {
					if !slices.Contains(graph[s], v) {
						connected = false
						break
					}
				}
				if connected {
					components[u] = append(components[u], v)
				}
			}
		}
	}
	p2 := 0
	var res string
	for _, v := range components {
		if len(v) > p2 {
			p2 = len(v)
			sort.Strings(v)
			res = strings.Join(v, ",")
		}
	}
	return fmt.Sprintf("%d", p1), fmt.Sprintf("%s", res)
}
