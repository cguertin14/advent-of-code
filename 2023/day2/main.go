package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	input string
	max   = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

func main() {
	lines := strings.Split(input, "\n")
	sum := 0
	sumPowers := 0
	for i, line := range lines[:len(lines)-1] {
		game := make(map[string]int)
		gameLine := strings.Split(line, ": ")
		splitted := strings.Split(gameLine[1], "; ")
		possible := true
		for _, s := range splitted {
			items := strings.Split(s, ", ")
			for _, item := range items {
				details := strings.Split(item, " ")
				count, _ := strconv.Atoi(details[0])
				color := details[1]
				if count > max[color] {
					possible = false
				}
				if game[color] < count {
					game[color] = count
				}
			}
		}
		if possible {
			sum += i + 1 // + 1 because indexes start at 1
		}
		power := 1
		for _, count := range game {
			power *= count
		}
		sumPowers += power
	}
	fmt.Println(sum)       // part 1
	fmt.Println(sumPowers) // part 2
}
