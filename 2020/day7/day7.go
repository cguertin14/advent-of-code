package day7

import (
	"fmt"

	"github.com/cguertin14/advent-of-code-2020/utils"
)

func main() {
	reader := utils.Reader{}
	_, err := reader.ReadFromFile("cmd/day7/input.txt")
	if err != nil {
		return
	}

	containers, containedBy := parseBags(reader)
	fmt.Printf("Part one: %d\n", partOne(containers, containedBy))
	fmt.Printf("Part two: %d\n", partTwo(containers))

	return
}

func partOne(containers map[string]map[string]int, containedBy map[string][]string) int {
	return countShinyGoldContainers(containers, containedBy)
}

func partTwo(containers map[string]map[string]int) int {
	return countShinyGoldContained(containers)
}
