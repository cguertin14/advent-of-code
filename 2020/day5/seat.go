package day5

import "github.com/cguertin14/advent-of-code-2020/utils"

// Seat is a struct containing a row and a column
type Seat struct {
	Row, Col int
}

// ID returns the ID of a Seat
func (seat *Seat) ID() int {
	return seat.Row*8 + seat.Col
}

func parseSeats(reader utils.Reader) (seats []Seat) {
	seats = make([]Seat, len(reader.Lines))

	for n, line := range reader.Lines {
		row := 0
		for i := 6; i >= 0; i-- {
			if line[6-i] == 'B' {
				row += 1 << i
			}
		}

		col := 0
		for i := 2; i >= 0; i-- {
			if line[9-i] == 'R' {
				col += 1 << i
			}
		}

		seats[n] = Seat{row, col}
	}

	return
}
