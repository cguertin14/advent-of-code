package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const (
	theirRockShape     string = "A"
	theirPaperShape    string = "B"
	theirScissorsShape string = "C"

	myRockShape     string = "X"
	myPaperShape    string = "Y"
	myScissorsShape string = "Z"

	needToWin  string = myScissorsShape
	needToDraw string = myPaperShape
	needToLose string = myRockShape
)

const (
	rockPoints int = 1 + iota
	paperPoints
	scissorsPoints

	drawPoints int = 3
	winPoints  int = 6
)

func main() {
	lines := strings.Split(input, "\n")

	scores1 := make([]int, len(lines))
	scores2 := make([]int, len(lines))
	for n, line := range lines {
		splitted := strings.Split(line, " ")
		opponentsMove := splitted[0]
		myMove := splitted[1]

		myScore1 := findOutcomeOfRound(myMove, opponentsMove)
		myScore2 := findOpiniatedOutcomeOfRound(myMove, opponentsMove)
		scores1[n] = myScore1
		scores2[n] = myScore2
	}

	printTotalScore(scores1)
	printTotalScore(scores2)
}

func printTotalScore(scores []int) {
	totalScore := 0
	for _, pts := range scores {
		totalScore += pts
	}
	fmt.Println(totalScore)
}

// part 2
func findOpiniatedOutcomeOfRound(whatINeedToDo, opponentsMove string) (myScore int) {
	myMove := ""

	switch whatINeedToDo {
	case needToWin:
		if opponentsMove == theirRockShape {
			myMove = myPaperShape
		} else if opponentsMove == theirPaperShape {
			myMove = myScissorsShape
		} else if opponentsMove == theirScissorsShape {
			myMove = myRockShape
		}
	case needToLose:
		if opponentsMove == theirRockShape {
			myMove = myScissorsShape
		} else if opponentsMove == theirPaperShape {
			myMove = myRockShape
		} else if opponentsMove == theirScissorsShape {
			myMove = myPaperShape
		}
	case needToDraw:
		myMove = opponentsMove
	}

	return findOutcomeOfRound(myMove, opponentsMove)
}

// part 1
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
	case theirRockShape, myRockShape:
		return rockPoints
	// Paper calculations
	case theirPaperShape, myPaperShape:
		return paperPoints
	// Scissors calculations
	case theirScissorsShape, myScissorsShape:
		return scissorsPoints
	default:
		return 0
	}
}
