package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	// part 1
	fmt.Printf("Part 1: %d\n", calculateMuls(input))

	// part 2
	part2Total := 0
	reg := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	matches := reg.FindAllString(input, -1)
	enabled := true
	for _, match := range matches {
		if match == "do()" {
			enabled = true
		} else if match == "don't()" {
			enabled = false
		} else if enabled {
			part2Total += calculateMuls(match)
		}
	}
	fmt.Printf("Part 2: %d\n", part2Total)
}

func calculateMuls(in string) (total int) {
	total = 0
	reg := regexp.MustCompile("(mul\\(\\d+,\\d+\\))")
	matches := reg.FindAllString(in, -1)
	for _, match := range matches {
		matchSub := match[3:]
		splitted := strings.Split(matchSub, ",")
		left, _ := strconv.Atoi(splitted[0][1:])
		right, _ := strconv.Atoi(splitted[1][:len(splitted[1])-1])
		total += (left * right)
	}
	return total
}
