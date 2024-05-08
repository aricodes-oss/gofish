package codebox

import (
	"fmt"
	"gofish/interpreter/parser"
	"os"
)

type Codebox struct {
	grid [][]parser.Token
}

// Creates a new, empty codesbox.
func NewCodebox() *Codebox {
	return &Codebox{}
}

// Reads a file into a codebox
func (cb *Codebox) BuildCodebox(filepath string) {

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := parser.NewScanner(file)

	// Clear codesbox, in case it already had data in it for some unfathomable reason
	cb.grid = make([][]parser.Token, 1)

	// keep track of how many rows and columns in the filepath
	// This is needed to make sure we pad the slice with appropriate whitespace
	x, y, local_x := 0, 0, 0

	// Read buffer token by token
	tk, _ := scanner.Scan()
	i := 0
	for tk != parser.EOF && i < 10 {
		switch tk {
		case parser.EOL:
			local_x = 0
			y++
		default:
			// Keep track of the longest line for padding later
			local_x++
			if local_x > x {
				x = local_x
			}

			// Store the token in the codebox grid
			if y > len(cb.grid)-1 {
				// if there are characters on a new line, add a new row.
				cb.addRow()
			}
			cb.grid[y] = append(cb.grid[y], tk)
		}

		tk, _ = scanner.Scan()
		i++

	}

	cb.padWhitespace(x)
	fmt.Println(cb.grid)
}

// Process the codebox so all rows are the same length
func (cb *Codebox) padWhitespace(to int) {
	for row := range cb.grid {
		for colSize := len(cb.grid[row]); colSize < to; colSize++ {
			cb.grid[row] = append(cb.grid[row], parser.WHITESPACE)
		}
	}
}

func (cb *Codebox) addRow() {
	cb.grid = append(cb.grid, make([]parser.Token, 0))
}
