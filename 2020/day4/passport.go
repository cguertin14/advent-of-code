package day4

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/cguertin14/advent-of-code-2020/utils"
)

var hexColor = regexp.MustCompile("^#[0-9a-f]{6}$")
var pidDigits = regexp.MustCompile("^[0-9]{9}$")

// Passport is a struct containing
// passport information
type Passport map[string]string

func readPassports(reader utils.Reader) (passports []Passport) {
	passports = make([]Passport, 0)
	for _, block := range reader.Lines {
		passport := make(Passport)
		for _, line := range strings.Split(block, "\n") {
			passport.build(line)
		}
		passports = append(passports, passport)
	}

	return
}

func (pass Passport) build(line string) {
	elements := strings.Split(line, " ")

	for _, element := range elements {
		line := strings.Split(element, ":")
		for i, value := range line {
			if i%2 != 0 {
				key := line[i-1]
				pass[key] = value
			}
		}
	}
}

// IsValid1 validates a passport for part one
func (pass Passport) IsValid1() bool {
	for _, val := range []string{"ecl", "eyr", "byr", "hcl", "pid", "hgt", "iyr"} {
		if pass[val] == "" {
			return false
		}
	}
	return true
}

// IsValid2 validates a passport for part two
func (pass Passport) IsValid2() bool {
	byr, err := strconv.Atoi(pass["byr"])
	if err != nil || byr < 1920 || byr > 2002 {
		return false
	}
	iyr, err := strconv.Atoi(pass["iyr"])
	if err != nil || iyr < 2010 || iyr > 2020 {
		return false
	}
	eyr, err := strconv.Atoi(pass["eyr"])
	if err != nil || eyr < 2020 || eyr > 2030 {
		return false
	}
	if len(pass["hgt"]) < 3 {
		return false
	}
	hgt, err := strconv.Atoi(pass["hgt"][:len(pass["hgt"])-2])
	units := pass["hgt"][len(pass["hgt"])-2:]
	switch units {
	case "cm":
		if hgt < 150 || hgt > 193 {
			return false
		}
		break
	case "in":
		if hgt < 59 || hgt > 76 {
			return false
		}
		break
	default:
		return false
	}
	if !hexColor.MatchString(pass["hcl"]) {
		return false
	}
	switch pass["ecl"] {
	case "amb":
	case "blu":
	case "brn":
	case "gry":
	case "grn":
	case "hzl":
	case "oth":
	default:
		return false
	}
	if !pidDigits.MatchString(pass["pid"]) {
		return false
	}
	return true
}
