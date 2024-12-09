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

// I think I actually learned a lot through this misinterpretation -- two lessons:
// 1) I got better at complicated two pointer shifting, I find these kinds of puzzles fun.
// 2) This has been a recurring theme recently, and one that I hope this will hammer home;
//    make sure I get the spec right!

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
	s1 := 0
	for id, pos := range ints {
		s1 += (id * pos)
	}

	return fmt.Sprintf("%d", s1)
}

func findFreeSpace(ints []T, j, n int) int {
	for i := range j {
		if ints[i].A == -1 && ints[i].B >= n {
			return i
		}
		i++
	}
	return -1
}

// Part 2
func day9(input []byte) (string, string) {
	lines := bytes.Split(input, []byte{'\n'})[0]
	linesCopy := make([]byte, len(lines))
	copy(linesCopy, lines)
	p1 := part1(linesCopy)
	ints := make([]T, 0)
	id := 0
	// Tuple is (id, count), where id = -1 for '.' positions.
	for i := range len(lines) {
		n := int(lines[i] - '0')
		if i%2 == 0 {
			ints = append(ints, T{id, n})
			id += 1
		} else {
			ints = append(ints, T{-1, n})
		}
	}
	j := len(ints) - 1
	for j > 1 {
		id, count := ints[j].A, ints[j].B
		if id == -1 {
			j -= 1
		} else {
			if i := findFreeSpace(ints, j, count); i != -1 {
				// If space is an exact match, we just overwrite the id.
				// Make sure to free up space we are swapping out of.
				if count == ints[i].B {
					ints[i].A = id
					ints[j].A = -1
					j -= 1
					// Otherwise we take some space and split it.
				} else {
					ints[j].A = -1
					leftover := T{-1, ints[i].B - count}
					ints = append(ints, T{-99, -99})
					copy(ints[i+2:], ints[i+1:])
					ints[i+1] = leftover
					ints[i].A = id
					ints[i].B = count
				}
			} else {
				j -= 1
			}
		}
	}
	pos := 0
	p2 := 0
	for i := range len(ints) {
		id, n := ints[i].A, ints[i].B
		if id != -1 {
			for range n {
				p2 += (id * pos)
				pos += 1
			}
		} else {
			pos += n
		}
	}

	return p1, fmt.Sprintf("%d", p2)
}
