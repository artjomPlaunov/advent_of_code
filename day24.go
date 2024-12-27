// day23.go
package main

import (
	"bytes"
	"fmt"
)

func day24(input []byte) (string, string) {
	parts := bytes.Split(input, []byte("\n\n"))
	fmt.Println(string(parts[0]))
	p1 := 0

	return fmt.Sprintf("%d", p1), ""
}
