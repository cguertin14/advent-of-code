package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	// Part 1
	lines := strings.Split(input, "\n")
	var gamma, epsilon string

	columns := make(map[int]map[string]int)
	for _, line := range lines {
		digits := strings.Split(line, "")
		for n, digit := range digits {
			if columns[n] == nil {
				columns[n] = make(map[string]int)
			}
			columns[n][digit]++
		}
	}

	keys := make([]int, 0)
	for key := range columns {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, k := range keys {
		values := columns[k]
		if values["1"] > values["0"] {
			gamma += "1"
			epsilon += "0"
			continue
		}
		gamma += "0"
		epsilon += "1"
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Printf("Answer 1: %d\n", gammaInt*epsilonInt)

	// Part 2
	var oxygen, scrubber int
	numbersToKeep := make([]string, 0)
	for n, char := range gamma {
		// 'lines' needs to change after each
		// iteration of 'most'.
		for _, line := range lines {
			for o, digit := range line {
				if n == o && digit == char {
					numbersToKeep = append(numbersToKeep, line)
				}
			}
		}

		// Chaque fois que 'lines' change,
		// on veut aussi que 'most' soit
		// mis à jour.
		// fmt.Println(position, len(most))

		lines = numbersToKeep
		numbersToKeep = make([]string, 0)
		break
	}

	fmt.Printf("Answer 2: %d\n", oxygen*scrubber)
}
