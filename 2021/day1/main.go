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

	var largerCount int32
	var prev string

	for _, line := range lines {
		if prev != "" {
			nInt, _ := strconv.Atoi(line)
			prevInt, _ := strconv.Atoi(prev)

			if nInt > prevInt {
				largerCount++
			}
		}
		prev = line
	}

	fmt.Printf("Answer 1: %d\n", largerCount)
}

func partTwo() {
	lines := strings.Split(input2, "\n")

	var largerCount int32
	var prevCount int

	for n, line := range lines {
		// Éviter les deux dernières lignes
		if n < len(lines)-2 {
			// Pour chaque ligne, additionner la ligne + les 2 suivantes
			current, _ := strconv.Atoi(line)
			next, _ := strconv.Atoi(lines[n+1])
			nextNext, _ := strconv.Atoi(lines[n+2])

			currentCount := current + next + nextNext
			if prevCount != 0 && currentCount > prevCount {
				largerCount++
			}

			prevCount = currentCount
		}
	}

	fmt.Printf("Answer 2: %d\n", largerCount)
}
