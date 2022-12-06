package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	var cratesA, cratesB crates
	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, "[") {
			for i := 1; i <= len(line); i += 4 {
				if line[i] != ' ' {
					cratesA = cratesA.insertAtBottom(i/4, line[i])
					cratesB = cratesB.insertAtBottom(i/4, line[i])
				}
			}
		}
		if line[0:4] == "move" {
			splitted := strings.Split(line, " ")
			origin, _ := strconv.Atoi(splitted[1])
			from, _ := strconv.Atoi(splitted[3])
			to, _ := strconv.Atoi(splitted[5])
			for i := 0; i < origin; i++ {
				cratesA.move(from-1, to-1, 1)
			}
			cratesB.move(from-1, to-1, origin)
		}
	}

	cratesA.printTop()
	cratesB.printTop()
}

type crates [][]byte

func (c crates) move(from, to, origin int) {
	vals := c[from][len(c[from])-origin : len(c[from])]
	c[from] = c[from][:len(c[from])-origin]
	c[to] = append(c[to], vals...)
}

func (c crates) printTop() {
	for i := 0; i < len(c); i++ {
		if len(c[i]) == 0 {
			fmt.Printf(" ")
			continue
		}
		// print last character of line
		// from bytes, to readable char
		fmt.Printf("%c", c[i][len(c[i])-1])
	}
	fmt.Println()
}

func (c crates) insertAtBottom(col int, val byte) crates {
	if len(c) <= col {
		backing := make(crates, col+1)
		copy(backing, c)
		c = backing
	}
	if c[col] == nil {
		c[col] = make([]byte, 0)
	}
	c[col] = append([]byte{val}, c[col]...)
	return c
}
