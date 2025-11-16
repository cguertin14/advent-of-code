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
	part1Count := 0
	part2Count := 0
	for _, line := range lines {
		digits := strings.Split(line, " ")
		// Part 1
		if testReport(digits) {
			part1Count++
			part2Count++
			continue
		}
		// Part 2
		for count := range digits {
			newDigits := make([]string, 0)
			for i, digit := range digits {
				if i != count {
					newDigits = append(newDigits, digit)
				}
			}
			if testReport(newDigits) {
				part2Count++
				break
			}
		}
	}
	fmt.Printf("Part 1: %d\n", part1Count)
	fmt.Printf("Part 2: %d\n", part2Count)
}

func testReport(digits []string) bool {
	prevDigit, _ := strconv.Atoi(digits[0])
	increasing := true
	for i, digit := range digits[1:] {
		currentDigit, _ := strconv.Atoi(digit)
		// check which way we are going, up or down at second digit.
		if i == 0 {
			increasing = currentDigit > prevDigit
		}
		prevDigit, _ = strconv.Atoi(digits[i])
		// check diff between numbers
		diff := prevDigit - currentDigit
		if increasing {
			diff = currentDigit - prevDigit
		}
		if diff < 1 || diff > 3 {
			return false // report is unsafe.
		}
		if (increasing && currentDigit < prevDigit) || (!increasing && currentDigit > prevDigit) {
			return false // unsafe is order changes.
		}
	}
	return true
}
