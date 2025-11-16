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
	safeCount := 0

outer:
	for _, line := range lines {
		digits := strings.Split(line, " ")
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
				continue outer // report is unsafe.
			}
			if (increasing && currentDigit < prevDigit) || (!increasing && currentDigit > prevDigit) {
				continue outer // unsafe is order changes.
			}
		}
		safeCount++
	}
	fmt.Printf("Part 1: %d\n", safeCount)
}
