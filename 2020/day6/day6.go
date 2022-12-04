package day6

import (
	"fmt"

	"github.com/cguertin14/advent-of-code-2020/utils"
)

func main() {
	reader := utils.Reader{}
	_, err := reader.ReadFromFile("cmd/day6/input.txt")
	if err != nil {
		return
	}

	groups, groupSize := parseGroups(reader)
	fmt.Printf("Part one: %d\n", partOne(groups))
	fmt.Printf("Part two: %d\n", partTwo(groups, groupSize))
}

func partOne(groups []Group) (sum int) {
	for _, group := range groups {
		sum += len(group)
	}
	return
}

func partTwo(groups []Group, groupSize []int) (sum int) {
	for n, group := range groups {
		sum += group.CountEveryone(n, groupSize)
	}
	return
}
