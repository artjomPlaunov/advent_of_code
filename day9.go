// day2.go
package main

import (
	"bytes"
	"fmt"
)

// func getDigits(num int) []int {
// 	if num == -1 {
// 		return []int{-1}
// 	}
// 	if num == 0 {
// 		return []int{0}
// 	}
// 	var digits []int
// 	for num > 0 {
// 		digits = append(digits, num%10)
// 		num /= 10
// 	}
// 	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
// 		digits[i], digits[j] = digits[j], digits[i]
// 	}
// 	return digits
// }

// I completely misinterpreted the problem statement for part1, and spent forever on it.
// I'm keeping this here since I think I did something much more complicated in terms of
// a 2 pointer algorithm, and it took quite a lot of work.
// The misinterpretation was assuming each blank space can only hold a single digit.

// Out gigabrained myself for pt 1.

// func day9(input []byte) (string, string) {
// 	lines := bytes.Split(input, []byte{'\n'})[0]
// 	ints := make([]int, 0)
// 	id := 0
// 	for i := range len(lines) {
// 		if i%2 == 0 {
// 			for j := 0; j < int(lines[i]-'0'); j++ {
// 				ints = append(ints, id)
// 			}
// 			id += 1
// 		} else {
// 			for j := 0; j < int(lines[i]-'0'); j++ {
// 				ints = append(ints, -1)
// 			}
// 		}
// 	}
// 	i := 0
// 	fmt.Println(ints)
// 	idMap := make(map[int]int)
// 	for i < len(ints) {
// 		if ints[len(ints)-1] == -1 {
// 			for ints[len(ints)-1] == -1 {
// 				ints = ints[:len(ints)-1]
// 			}
// 			continue
// 		}
// 		if ints[i] != -1 {
// 			idMap[i] = ints[i]
// 			i += 1
// 		} else {
// 			n := ints[len(ints)-1]
// 			ints = ints[:len(ints)-1]
// 			for _, d := range getDigits(n) {
// 				for i < len(ints) && ints[i] != -1 {
// 					idMap[i] = ints[i]
// 					i += 1
// 				}
// 				if i < len(ints) {
// 					ints[i] = d
// 					idMap[i] = n
// 					i += 1
// 				} else {
// 					ints = append(ints, d)
// 					idMap[i] = n
// 					i += 1
// 				}
// 			}
// 		}
// 	}
// 	keys := make([]int, 0, len(idMap))
// 	for k := range idMap {
// 		keys = append(keys, k)
// 	}
// 	fmt.Println(ints)
// 	sort.Ints(keys)
// 	pos := 0
// 	s1 := 0
// 	for _, k := range keys {
// 		n := idMap[k]
// 		for range len(getDigits(n)) {
// 			s1 += pos * n
// 			pos += 1
// 		}
// 	}

// 	return fmt.Sprintf("%d", s1), ""
// }

func part1(input []byte) string {
	lines := bytes.Split(input, []byte{'\n'})[0]
	ints := make([]int, 0)
	id := 0
	for i := range len(lines) {
		if i%2 == 0 {
			for j := 0; j < int(lines[i]-'0'); j++ {
				ints = append(ints, id)
			}
			id += 1
		} else {
			for j := 0; j < int(lines[i]-'0'); j++ {
				ints = append(ints, -1)
			}
		}
	}
	i := 0
	idMap := make(map[int]int)
	for i < len(ints) {
		if ints[len(ints)-1] == -1 {
			for ints[len(ints)-1] == -1 {
				ints = ints[:len(ints)-1]
			}
			continue
		}
		if ints[i] != -1 {
			idMap[i] = ints[i]
			i += 1
		} else {
			n := ints[len(ints)-1]
			ints = ints[:len(ints)-1]
			ints[i] = n
			idMap[i] = n
			i += 1
		}
	}
	fmt.Println(ints)
	s1 := 0
	for id, pos := range ints {
		s1 += (id * pos)
	}

	return fmt.Sprintf("%d", s1)
}

func day9(input []byte) (string, string) {
	lines := bytes.Split(input, []byte{'\n'})[0]
	linesCopy := make([]byte, len(lines))
	copy(linesCopy, lines)
	p1 := part1(linesCopy)
	return p1, ""
}
