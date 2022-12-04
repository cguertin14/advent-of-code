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

	totals := make([]int, len(calories))
	for n, arr := range calories {
		for _, calorie := range arr {
			totals[n] += calorie
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(totals)))

	fmt.Println(findHighestCalories(totals))
	fmt.Println(top3Totals(totals))
}

func findHighestCalories(totals []int) int {
	return totals[0]
}

func top3Totals(totals []int) int {
	return totals[0] + totals[1] + totals[2]
}
