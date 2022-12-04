package day4

import (
	"fmt"

	"github.com/cguertin14/advent-of-code-2020/utils"
)

func main() {
	reader := utils.Reader{}
	_, err := reader.ReadFromFileDouble("cmd/day4/input.txt")
	if err != nil {
		return
	}

	passports := readPassports(reader)
	fmt.Printf("Part One: %d \n", partOne(passports))
	fmt.Printf("Part Two: %d \n", partTwo(passports))

	return
}

func partOne(passports []Passport) (count int) {
	for _, pass := range passports {
		if pass.IsValid1() {
			count++
		}
	}
	return
}

func partTwo(passports []Passport) (count int) {
	for _, pass := range passports {
		if pass.IsValid2() {
			count++
		}
	}
	return
}
