// day1.go
package main

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"
)

func day1(input []byte) (string, string) {
	var list1, list2 []int
	lines := bytes.Split(input, []byte{'\n'})
	for _, line := range lines {
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}
		parts := bytes.Fields(line)
		num1, _ := strconv.Atoi(string(parts[0]))
		num2, _ := strconv.Atoi(string(parts[1]))
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	var distance int
	sort.Ints(list1)
	sort.Ints(list2)
	for i := 0; i < len(list1); i++ {
		n1 := list1[i]
		n2 := list2[i]
		distance += int(math.Abs(float64(n1) - float64(n2)))
	}
	freq1 := make(map[int]int)
	freq2 := make(map[int]int)
	for _, num := range list1 {
		freq1[num] += 1
	}
	for _, num := range list2 {
		freq2[num] += 1
	}
	similarity := 0
	for _, k := range list1 {
		similarity += (k * freq2[k])
	}

	return fmt.Sprintf("%d", distance), fmt.Sprintf("%d", similarity)

}
