package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(findMarker(3))
	fmt.Println(findMarker(13))
}

func findMarker(minCharPosition int) (marker int) {
outer:
	for i, char := range input {
		if i >= minCharPosition {
			chars := map[byte]bool{byte(char): true}
			for j := i - 1; j > i-(minCharPosition+1); j-- {
				if _, ok := chars[input[j]]; ok {
					continue outer
				}
				chars[input[j]] = true
			}
			marker = i + 1
			break outer
		}
	}
	return
}
