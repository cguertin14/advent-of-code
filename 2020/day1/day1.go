package day1

import (
	"fmt"

	"github.com/cguertin14/advent-of-code-2020/utils"
)

func main() {
	reader := utils.Reader{}
	_, err := reader.ReadFromFile("cmd/day1/input.txt")
	if err != nil {
		return
	}

	// Get results of each challenge.
	numbers := reader.ToNumbers()
	one := partOne(numbers)
	two := partTwo(numbers)

	// SHOW RESULTS
	fmt.Println(fmt.Sprintf("Part one, the answer is: %d", one))
	fmt.Println(fmt.Sprintf("Part two, the answer is: %d", two))
}

func partOne(numbers []int) int {
	for _, num := range numbers {
		for _, otherNum := range numbers {
			if num+otherNum == 2020 {
				return num * otherNum
			}
		}
	}

	return 0
}

func partTwo(numbers []int) int {
	for _, num := range numbers {
		for _, secondNum := range numbers {
			for _, thirdNum := range numbers {
				if num+secondNum+thirdNum == 2020 {
					return num * secondNum * thirdNum
				}
			}
		}
	}

	return 0
}
