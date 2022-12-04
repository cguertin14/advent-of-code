package day3

import (
	"fmt"

	"github.com/cguertin14/advent-of-code-2020/utils"
)

// Down represent the move length, downwards
const Down = 1

// Right represent the move length, to the right
const Right = 3

// Execute runs the command.
func main() {
	reader := utils.Reader{}
	_, err := reader.ReadFromFile("cmd/day3/input.txt")
	if err != nil {
		return
	}

	fmt.Printf("Part one: %d\n", partOne(reader))
	fmt.Printf("Part two: %d\n", partTwo(reader))
}

func countTrees(reader utils.Reader) (trees [][]bool) {
	lines := reader.Lines
	trees = make([][]bool, len(lines))

	for i, s := range lines {
		trees[i] = make([]bool, len(s))
		for j, c := range s {
			trees[i][j] = (c == '#')
		}
	}

	return
}

func checkSlope(down, right int, trees [][]bool) (count int) {
	count = 0
	for index := 0; index*down < len(trees); index++ {
		current := index * down
		col := (index * right) % len(trees[current])
		if trees[current][col] {
			count++
		}
	}
	return
}

func partOne(reader utils.Reader) int {
	return checkSlope(1, 3, countTrees(reader))
}

func partTwo(reader utils.Reader) int {
	trees := countTrees(reader)

	a := checkSlope(1, 1, trees)
	b := checkSlope(1, 3, trees)
	c := checkSlope(1, 5, trees)
	d := checkSlope(1, 7, trees)
	e := checkSlope(2, 1, trees)

	return a * b * c * d * e
}
