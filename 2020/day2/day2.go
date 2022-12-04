package day2

import (
	"fmt"

	"github.com/cguertin14/advent-of-code-2020/utils"
)

func main() {
	reader := utils.Reader{}
	_, err := reader.ReadFromFile("cmd/day2/input.txt")
	if err != nil {
		return
	}

	// 1- Build all passwords.
	passwords := buildPasswords(reader)

	// 2- Validate all passwords (part 1)
	validCount := validatePasswordsA(passwords)

	// 3- Print the total of valid passwords.
	fmt.Println(fmt.Sprintf("There are %d valid passwords in the list. (part 1)", validCount))

	// 4- Validate all passwords (part 2)
	validCount = validatePasswordsB(passwords)

	// 3- Print the total of valid passwords.
	fmt.Println(fmt.Sprintf("There are %d valid passwords in the list. (part 2)", validCount))
}
