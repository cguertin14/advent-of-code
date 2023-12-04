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
)

func main() {
	lines := strings.Split(input, "\n")
	points := 0
	tree := make(map[int][]int)
	queue := make(map[int]int)

	for i, line := range lines[:len(lines)-1] {
		numbers := strings.Split(line, ": ")[1]
		table := strings.Split(numbers, " | ")
		winningNumbers := buildSliceFromStrSlice(strings.Split(table[0], " "))
		myNumbers := buildSliceFromStrSlice(strings.Split(table[1], " "))
		queue[i] = 1

		lastPower := 0
		winners := 0
		for _, n := range myNumbers {
			for _, o := range winningNumbers {
				if n == o {
					winners++
					tree[i] = append(tree[i], i+winners)
					if lastPower == 0 {
						lastPower = 1
						continue
					}
					lastPower *= 2
				}
			}
		}
		points += lastPower
	}
	fmt.Println(points) // part 1

	// part 2
	inventory := make(map[int]int)
	for len(queue) != 0 {
		next := make(map[int]int)
		for k, v := range queue {
			inventory[k] += v
			for _, a := range tree[k] {
				next[a] += v
			}
		}
		queue = next
	}
	var cards int
	for _, v := range inventory {
		cards += v
	}
	fmt.Println(cards) // part 2 result
}

func buildSliceFromStrSlice(s []string) []int {
	out := make([]int, 0)
	for _, str := range s {
		if strings.TrimSpace(str) != "" {
			n, _ := strconv.Atoi(str)
			out = append(out, n)
		}
	}
	return out
}
