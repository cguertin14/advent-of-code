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

type coord struct {
	x, y int
}

type symbol struct {
	char        rune
	coordinates coord
}

func main() {
	// Any number adjacent to a symbol, even diagonally, is
	// a "part number" and should be included in your sum.
	//
	// symbol = char != "."
	//

	// how to solve this?
	// -> build a map of the game to search into.
	// -> each number must search the previous, current (left-right)
	//    and next line for adjacent (directly under/over or left/right)
	// -> to do so, first line can't check previous one,
	//    and last line can't check the next line.
	//
	// how to build a map of the board?
	// -> list of lines should do. can then compare each char with
	//    the current index to check adjacent symbols.
	// -> we need to identify the numbers on each line (i.e: 523)
	//    so that we can compare them in real time with prev/next line.

	lines := strings.Split(input, "\n")
	sum, sum2 := 0, 0
	gears := make(map[coord][]int)
	for i, line := range lines[:len(lines)-1] {
		newIndex := 0
		for j, char := range line {
			// check here that current number hasn't been checked already
			if newIndex == 0 || j == newIndex {
				newIndex = 0
				if _, err := strconv.Atoi(string(char)); err == nil {
					// char is a digit, now fetch all the other digits next to it.
					currentNumber := string(char)
					for k, otherChar := range line[j+1:] {
						if _, err := strconv.Atoi(string(otherChar)); err == nil {
							// otherChar is a digit, append to currentNumber
							currentNumber += string(otherChar)
							// what is the goal of 'newIndex'?
							// -> the goal is to identify unique numbers once.
							// -> its goal is to set a new index for the next
							//    iteration around 'j := range line'.
							// -> the nextIndex sets the position where the next iteration will take place.
							newIndex = k + j + 2
						} else {
							break
						}
					}

					firstLineBuffer := 0
					linesToCheck := make([]string, 0)
					if i == 0 {
						// only do current & next lines
						linesToCheck = []string{line, lines[i+1]}
						firstLineBuffer = i
					} else if i == len(lines)-2 { // -2 to skip last empty line.
						// only do current & previous lines
						linesToCheck = []string{lines[i-1], line}
						firstLineBuffer = i - 1
					} else {
						// do all lines otherwise.
						linesToCheck = []string{lines[i-1], line, lines[i+1]}
						firstLineBuffer = i - 1
					}

					// perform checks on previous, current & next line.
					isIt, s := isSymbolAdjacent(line, linesToCheck, j, j+len(currentNumber), firstLineBuffer)
					number, _ := strconv.Atoi(currentNumber)
					if isIt {
						sum += number
					}
					if s.char == '*' {
						gears[s.coordinates] = append(gears[s.coordinates], number)
					}

					// reset currentNumber
					currentNumber = ""
				}
			}
		}

	}
	fmt.Println(sum) // part 1

	for _, gear := range gears {
		if len(gear) == 2 {
			sum2 += gear[0] * gear[1]
		}
	}
	fmt.Println(sum2) // part 2
}

func isSymbolAdjacent(
	line string,
	linesToCheck []string,
	startIndexInLine, endIndexInLine, firstLineBuffer int,
) (adjacent bool, s symbol) {
	// how do you do checks?
	// -> first, gotta identify wether to do the previous line,
	//    that is if the line is not the first one
	// -> second, identify wether to do next line, that is if
	//    we are not on the last line currently.
	// -> third, identify range of checks. that would be the
	//    position of the first digit - 1, up to the position
	// 		of the last digit + 1.
	// -> then, check the previous lines to see a symbol (not '.')
	// -> if one is found in the range, currentNumber is a part number.
	// -> add currentNumber to sum.

	// identify range.
	indexToEnd := endIndexInLine + 1
	if endIndexInLine == len(line) {
		indexToEnd = endIndexInLine
	}
	indexToStart := startIndexInLine - 1
	if startIndexInLine == 0 {
		indexToStart = startIndexInLine
	}

	for i, lineToCheck := range linesToCheck {
		for j, char := range lineToCheck[indexToStart:indexToEnd] {
			// if it's a number, skip
			if _, err := strconv.Atoi(string(char)); err == nil {
				continue
			}
			// else, check everything
			if string(char) != "." {
				return true, symbol{
					char: char,
					coordinates: coord{
						x: i + firstLineBuffer,
						y: j + indexToStart,
					},
				}
			}
		}
	}

	return
}
