package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input1.txt
var input1 string

//go:embed input2.txt
var input2 string

func main() {
	partOne()
	partTwo()
}

func partOne() {
	lines := strings.Split(input1, "\n")
	var horizontal, depth int

	for _, line := range lines {
		elems := strings.Split(line, " ")
		if len(elems) == 2 {
			movement, _ := strconv.Atoi(elems[1])
			switch elems[0] {
			case "up":
				depth -= movement
			case "down":
				depth += movement
			case "forward":
				horizontal += movement
			}
		}
	}
	fmt.Printf("Answer 1: %d\n", horizontal*depth)
}

func partTwo() {
	lines := strings.Split(input2, "\n")
	var horizontal, depth, aim int

	for _, line := range lines {
		elems := strings.Split(line, " ")
		if len(elems) == 2 {
			movement, _ := strconv.Atoi(elems[1])
			switch elems[0] {
			case "up":
				aim -= movement
			case "down":
				aim += movement
			case "forward":
				horizontal += movement
				depth += aim * movement
			}
		}
	}
	fmt.Printf("Answer 2: %d\n", horizontal*depth)
}
