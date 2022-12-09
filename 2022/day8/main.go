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

	// build forest
	forest := make([][]int, len(lines))
	for i, line := range lines {
		for _, char := range line {
			intC, _ := strconv.Atoi(string(char))
			forest[i] = append(forest[i], intC)
		}
	}

	// outside visible trees that we can skip
	// - minus 4 corners to de-double trees
	visibleTrees := len(lines)*2 + len(lines[0])*2 - 4 // 16 initial trees in test
	highestScenicScore := 0
	for rowNumber := 1; rowNumber < len(forest[1:]); rowNumber++ {
		row := forest[rowNumber]
		for columnNumber := 1; columnNumber < len(row[1:]); columnNumber++ {
			tree := row[columnNumber]                               // current tree
			left := row[0:columnNumber]                             // trees left of current tree
			right := row[columnNumber+1:]                           // trees right of current tree
			up := buildValues(forest[0:rowNumber], columnNumber)    // trees up of current tree
			down := buildValues(forest[rowNumber+1:], columnNumber) // trees down of current tree

			// calculate scenic score of current tree and check if it is the highest
			score := calculateScore(tree, left, right, up, down)
			if score > highestScenicScore {
				highestScenicScore = score
			}

			// check rows
			if isTreeVisible(tree, left) || isTreeVisible(tree, right) {
				visibleTrees++
				continue
			}

			// check columns
			if isTreeVisible(tree, up) || isTreeVisible(tree, down) {
				visibleTrees++
			}
		}
	}
	// part 1
	fmt.Printf("%d\n", visibleTrees)
	// part 2
	fmt.Printf("%d\n", highestScenicScore)
}

func calculateScore(tree int, left, right, up, down []int) int {
	newLeft, newUp := make([]int, len(left)), make([]int, len(up))
	copy(newLeft, left)
	left = newLeft
	copy(newUp, up)
	up = newUp

	// reverse order of left and up trees
	reverseSlice(left)
	reverseSlice(up)

	calc := func(trees []int) (count int) {
		for _, other := range trees {
			if other >= tree {
				count++
				break
			}
			count++
		}
		return
	}
	scores := [4]int{calc(up), calc(left), calc(down), calc(right)}
	score := 1
	for _, s := range scores {
		score *= s
	}
	return score
}

func reverseSlice(slice []int) {
	sort.SliceStable(slice, func(i, j int) bool {
		return i > j
	})
}

func buildValues(rows [][]int, columnNumber int) []int {
	values := make([]int, len(rows))
	for i, val := range rows {
		values[i] = val[columnNumber]
	}
	return values
}

func isTreeVisible(tree int, otherTrees []int) bool {
	for _, otherTree := range otherTrees {
		if otherTree >= tree {
			return false
		}
	}
	return true
}
