package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const (
	rockPoints int = 1 + iota
	paperPoints
	scissorsPoints

	drawPoints int = 3
	winPoints  int = 6
)

func main() {
	lines := strings.Split(input, "\n")

	scores := make([]int, len(lines))
	for n, line := range lines {
		splitted := strings.Split(line, " ")
		opponentsMove := splitted[0]
		myMove := splitted[1]

		myScore := findOutcomeOfRound(myMove, opponentsMove)
		scores[n] = myScore
	}

	totalScore := 0
	for _, pts := range scores {
		totalScore += pts
	}
	fmt.Println(totalScore)
}

func findOutcomeOfRound(myMove, opponentsMove string) (myScore int) {
	myScore = calculateShapeScore(myMove)
	opponentsScore := calculateShapeScore(opponentsMove)

	// Check if draw
	if myScore == opponentsScore {
		myScore += drawPoints
	} else {
		// rock defeats scissors
		if myScore == rockPoints && opponentsScore == scissorsPoints ||
			// paper defeats rock
			myScore == paperPoints && opponentsScore == rockPoints ||
			// scissors defeats paper
			myScore == scissorsPoints && opponentsScore == paperPoints {
			// Take the W
			myScore += winPoints
		}
	}

	return
}

func calculateShapeScore(move string) int {
	switch move {
	// Rock calculations
	case "A", "X":
		return rockPoints
	// Paper calculations
	case "B", "Y":
		return paperPoints
	// Scissors calculations
	case "C", "Z":
		return scissorsPoints
	default:
		return 0
	}
}
