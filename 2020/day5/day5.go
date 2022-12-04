package day5

import (
	"fmt"

	"github.com/cguertin14/advent-of-code-2020/utils"
)

func main() {
	reader := utils.Reader{}
	_, err := reader.ReadFromFile("cmd/day5/input.txt")
	if err != nil {
		return
	}

	seats := parseSeats(reader)
	fmt.Printf("Part one: %d\n", partOne(seats))
	fmt.Printf("Part two: %d\n", partTwo(seats))
}

func partOne(seats []Seat) (highest int) {
	for _, seat := range seats {
		if seat.ID() > highest {
			highest = seat.ID()
		}
	}
	return
}

func partTwo(seats []Seat) int {
	traversed := make([]bool, 1024) // 2**10
	for _, s := range seats {
		traversed[s.ID()] = true
	}

	init := true
	for i, found := range traversed {
		if init && found {
			init = false
			continue
		}
		if !init && !found {
			return i
		}
	}

	return 0
}
