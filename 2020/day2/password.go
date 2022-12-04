package day2

import (
	"strconv"
	"strings"

	"github.com/cguertin14/advent-of-code-2020/utils"
)

// Password is a struct containing
// password-related data.
type Password struct {
	NumA    int
	NumB    int
	Letter  string
	Content string
}

// buildPasswords builds an array of passwords.
func buildPasswords(reader utils.Reader) (result []*Password) {
	for _, line := range reader.Lines {
		result = append(result, (&Password{}).Build(line))
	}
	return
}

// validatePasswordsA validates an array of passwords.
func validatePasswordsA(passwords []*Password) (total int) {
	for _, pwd := range passwords {
		if pwd.IsValidA() {
			total++
		}
	}
	return
}

// validatePasswordsB validates an array of passwords.
func validatePasswordsB(passwords []*Password) (total int) {
	for _, pwd := range passwords {
		if pwd.IsValidB() {
			total++
		}
	}
	return
}

// Build is a function to build the struct "Password"
// with its attributes.
func (pass *Password) Build(line string) *Password {
	elements := strings.Split(line, " ") // Returns an array with 3 elements in it.

	pattern := strings.Split(elements[0], "-")
	pass.NumA, _ = strconv.Atoi(pattern[0])
	pass.NumB, _ = strconv.Atoi(pattern[1])

	pass.Letter = strings.Replace(elements[1], ":", "", 1)
	pass.Content = elements[2]

	return pass
}

// IsValidA validates if a password is valid.
func (pass Password) IsValidA() bool {
	letterCount := 0

	for _, char := range pass.Content {
		if string(char) == pass.Letter {
			letterCount++
		}
	}

	return letterCount >= pass.NumA && letterCount <= pass.NumB
}

// IsValidB validates if a password is
// valid for the second problem.
func (pass Password) IsValidB() bool {
	letterCount := 0

	for pos, char := range pass.Content {
		if (pos == pass.NumA-1 || pos == pass.NumB-1) && string(char) == pass.Letter {
			letterCount++
		}
	}

	return letterCount == 1
}
