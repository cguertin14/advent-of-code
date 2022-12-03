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

	// The winner of the whole tournament is the player with the highest score.
	// Your total score is the sum of your scores for each round.
	// The score for a single round is the score for the shape you selected
	// (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for
	// the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

	// What would your total score be if everything goes exactly according to your strategy guide?
	// -> PRINT TOTAL SCORE.
	totalScore := 0
	for _, pts := range scores {
		totalScore += pts
	}
	fmt.Println(totalScore)
}

func findOutcomeOfRound(myMove, opponentsMove string) (myScore int) {
	// Rock defeats Scissors
	// Paper defeats Rock
	// Scissors defeats Paper

	// If both players choose the same shape, the round instead ends in a draw.

	// A for Rock, B for Paper, and C for Scissors -> Opponent's move
	// X for Rock, Y for Paper, and Z for Scissors -> My move

	// init values with shape scores
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
		// else {
		// 	// Take the L
		// 	opponentsScore += winPoints
		// }
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
