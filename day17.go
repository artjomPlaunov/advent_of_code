// day2.go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func backtrack(quine []int, ip, AA int) (bool, int) {
	for i := range 8 {
		A := (AA << 3) + i
		B := i
		B = B ^ 3
		C := A / (1 << B)
		B = B ^ C
		B = B ^ 5
		if (B % 8) == quine[ip] {
			if ip == 0 {
				return true, A
			} else {
				found, a := backtrack(quine, ip-1, A)
				if found {
					return found, a
				}
			}
		}
	}
	return false, 0
}

func day17(input []byte) (string, string) {
	// A := 729
	// B := 0
	// C := 0
	// program := []int{0, 1, 5, 4, 3, 0}
	// A := 117440
	// B := 0
	// C := 0
	// program := []int{0, 3, 5, 4, 3, 0}
	A := 236580836040301
	B := 0
	C := 0
	program := []int{2, 4, 1, 3, 7, 5, 0, 3, 4, 3, 1, 5, 5, 5, 3, 0}
	combo := func(op int) int {
		if 0 <= op && op <= 3 {
			return op
		}
		if op == 4 {
			return A
		}
		if op == 5 {
			return B
		}
		if op == 6 {
			return C
		}
		return -1
	}

	ip := 0
	output := make([]int, 0)
	for ip < len(program) {
		fmt.Println(A, B, C)
		if program[ip] == 0 {
			//adv
			A = A / (1 << combo(program[ip+1]))
			ip += 2
		} else if program[ip] == 1 {
			// bxl
			B = B ^ program[ip+1]
			ip += 2
		} else if program[ip] == 2 {
			// bst
			B = combo(program[ip+1]) % 8
			ip += 2
		} else if program[ip] == 3 {
			// jnz
			if A == 0 {
				ip += 2
			} else {
				ip = program[ip+1]
			}
		} else if program[ip] == 4 {
			// bxc
			B = B ^ C
			ip += 2
		} else if program[ip] == 5 {
			// out
			output = append(output, combo(program[ip+1])%8)
			ip += 2
		} else if program[ip] == 6 {
			// bdv
			B = A / (1 << combo(program[ip+1]))
			ip += 2
		} else if program[ip] == 7 {
			// cdv
			C = A / (1 << combo(program[ip+1]))
			ip += 2
		}
	}
	res := make([]string, 0)
	for _, n := range output {
		res = append(res, strconv.Itoa(n))
	}
	fmt.Println(backtrack(program, len(program)-1, 0))
	return strings.Join(res, ","), ""
}
