package day6

import (
	"github.com/cguertin14/advent-of-code-2020/utils"
)

// Group is a map of rune (letter) and int (count)
type Group map[rune]int

func parseGroups(reader utils.Reader) (groups []Group, groupSize []int) {
	groups = []Group{make(Group)}
	groupSize = []int{0}
	groupNum := 0

	for _, s := range reader.Lines {
		if s == "" {
			groups = append(groups, make(Group))
			groupSize = append(groupSize, 0)
			groupNum++
			continue
		}

		for _, char := range s {
			groups[groupNum][char]++
		}
		groupSize[groupNum]++
	}

	return
}

// CountEveryone counts only if every person
// of a group answered "yes" to the same questions.
func (group Group) CountEveryone(n int, groupSize []int) (total int) {
	for _, count := range group {
		if count == groupSize[n] {
			total++
		}
	}
	return
}
