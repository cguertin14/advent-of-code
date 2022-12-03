package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	notEmptyLines := 0
	prevLine := ""
	for _, line := range lines {
		if prevLine == "" {
			notEmptyLines++
		}
		prevLine = line
	}

	calories := make([][]int, notEmptyLines)
	elf := 0
	for _, line := range lines {
		// continue adding to elf inventory
		// while line != empty str
		if line == "" {
			elf++
			continue
		}

		calorie, _ := strconv.Atoi(line)
		calories[elf] = append(calories[elf], calorie)
	}

	println(findHighestCalories(calories))
}

func findHighestCalories(calories [][]int) int {
	totals := make([]int, len(calories))
	for n, arr := range calories {
		for _, calorie := range arr {
			totals[n] += calorie
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totals)))

	return totals[0]
}
