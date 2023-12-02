package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	input         string
	numbersMapped = map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}
)

func main() {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		var localNumber, lastNumber, stringNumber string
		for i, c := range line {
			// string validation - here, need to validate
			// the current char against all remaining chars.
			for _, r := range line[i:] {
				stringNumber += string(r)
				if val, ok := numbersMapped[stringNumber]; ok {
					strN := strconv.Itoa(val)
					if len(localNumber) == 0 { // set first number
						localNumber = strN
					}
					lastNumber = strN // keep track of last number
					continue
				}
			}
			stringNumber = ""                                  // reset string
			if _, err := strconv.Atoi(string(c)); err != nil { // int validation
				continue // skip current char if not an int.
			}
			if len(localNumber) == 0 { // set first number
				localNumber = string(c)
			}
			lastNumber = string(c) // keep track of last number
		}
		localNumber += lastNumber
		val, _ := strconv.Atoi(localNumber)
		sum += val
	}
	fmt.Println(sum)
}
