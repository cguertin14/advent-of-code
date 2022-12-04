package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Sections []int

func main() {
	lines := strings.Split(input, "\n")

	// List of pairs
	// each pair has 2 elves
	// each elf has many sections
	// [[elf1, elf2], [elf3, elf4], [elf5, elf6], ...]

	pairs := make([][]Sections, len(lines))
	pairsContained := 0
	pairsThatOverlap := 0
	for i, line := range lines {
		splitted := strings.Split(line, ",")
		section1 := countSections(splitted[0])
		section2 := countSections(splitted[1])
		pairs[i] = append(pairs[i], section1, section2)
		if doesSectionContainOtherSection(section1, section2) || doesSectionContainOtherSection(section2, section1) {
			pairsContained++
		}
		if doesSectionOverlapInOtherSection(section1, section2) || doesSectionOverlapInOtherSection(section2, section1) {
			pairsThatOverlap++
		}
	}

	fmt.Println(pairsContained)
	fmt.Println(pairsThatOverlap)
}

func doesSectionOverlapInOtherSection(a, b Sections) bool {
	exists := make(map[int]bool)
	for _, value := range a {
		exists[value] = true
	}
	for _, value := range b {
		if exists[value] {
			return true
		}
	}
	return false
}

func doesSectionContainOtherSection(a, b Sections) bool {
	exists := make(map[int]bool)
	for _, value := range a {
		exists[value] = true
	}
	for _, value := range b {
		if !exists[value] {
			return false
		}
	}
	return true
}

func countSections(section string) (sections Sections) {
	splitted := strings.Split(section, "-")
	begin := splitted[0]
	end := splitted[1]

	beginI, _ := strconv.Atoi(begin)
	endI, _ := strconv.Atoi(end)
	sections = make(Sections, endI-beginI+1)
	index := 0
	for i := beginI; i <= endI; i++ {
		sections[index] = i
		index++
	}

	return
}
