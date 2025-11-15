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
	left := make([]int64, len(lines))
	right := make([]int64, len(lines))
	for i, line := range lines {
		elems := strings.Split(line, "   ")
		first := elems[0]
		second := elems[1]
		left[i], _ = strconv.ParseInt(first, 10, 64)
		right[i], _ = strconv.ParseInt(second, 10, 64)
	}

	// Sort  the 2 slices
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	// calculate the sum...
	sum := int64(0)
	for i := range left {
		if right[i] > left[i] {
			sum += right[i] - left[i]
		} else {
			sum += left[i] - right[i]
		}
	}
	fmt.Printf("Part 1: %d\n", sum)

	// part 2: check how many times the left numbers appears in the right list,
	// multiply said number by number of times it appears in the other list.
	sum2 := int64(0)
	for i := range left {
		count := int64(0)
		for j := range right {
			if left[i] == right[j] {
				count++
			}
		}
		sum2 += left[i] * count
	}
	fmt.Printf("Part 2: %d\n", sum2)
}
