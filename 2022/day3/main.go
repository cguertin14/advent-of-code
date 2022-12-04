package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const (
	alphabetStr string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func main() {
	lines := strings.Split(input, "\n")
	total1 := 0
	groups := make([][]string, len(lines)/3)
	accumulatedSubGroup := make([]string, 0)
	subGroupIndex := 0

	for i, line := range lines {
		halfPoint := len(line) / 2
		firstCompartment := line[:halfPoint]
		secondCompartment := line[halfPoint:]

		chars := findCommonCharsInTwoStrings(firstCompartment, secondCompartment)
		localTotal1 := findTotalFromChars(chars)
		total1 += localTotal1

		// part 2
		accumulatedSubGroup = append(accumulatedSubGroup, line)
		if (i+1)%3 == 0 {
			groups[subGroupIndex] = accumulatedSubGroup
			accumulatedSubGroup = []string{}
			subGroupIndex++
		}
	}

	// part 2
	total2 := 0
	for _, subGroup := range groups {
		chars := make([]rune, 0)
		for _, str1 := range subGroup {
			for _, str2 := range subGroup {
				localChars := findCommonCharsInTwoStrings(str1, str2)
				chars = findCommonCharsInTwoStrings(string(localChars), string(chars))
				if len(chars) == 0 {
					chars = localChars
				}
			}
		}

		localTotal := findTotalFromChars(chars)
		total2 += localTotal
	}

	fmt.Println(total1)
	fmt.Println(total2)
}

func findTotalFromChars(chars []rune) (total int) {
	total = 0
	for _, char := range chars {
		total += findLetterIndexInAlphabet(char)
	}
	return
}

func arrayIncludesChar(char rune, chars []rune) bool {
	for _, c := range chars {
		if c == char {
			return true
		}
	}
	return false
}

func findCommonCharsInTwoStrings(a, b string) (c []rune) {
	// Build a reference table to lookup into
	// and set to true each char of 'a' str
	referenceTable := make(map[rune]bool)
	for _, item := range a {
		referenceTable[item] = true
	}
	// Lookup each char from 'b' str into
	// reference table
	for _, item := range b {
		// Make sure array does not already include char
		if _, ok := referenceTable[item]; ok && !arrayIncludesChar(item, c) {
			c = append(c, item)
		}
	}
	return
}

func findLetterIndexInAlphabet(letter rune) (res int) {
	for n, letterInAlphabet := range alphabetStr {
		if letter == letterInAlphabet {
			res = n + 1
			break
		}
	}
	return
}
