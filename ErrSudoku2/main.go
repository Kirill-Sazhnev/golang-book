package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	rows, columns = 9, 9
	empty         = 0
)

// Cell является квадратом сетки Судоку
type Cell struct {
	digit int8
	fixed bool
}

// Grid является сеткой Судоку
type Grid [rows][columns]Cell

var (
	ErrBounds  = errors.New("out of grid boundaries")
	ErrDigit   = errors.New("invalid digit")
	ErrInitial = errors.New("Initial digit cannot be changed")
)

func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return ErrBounds
	} else if g.initialDigit(row, column) {
		return ErrInitial
	}
	if !validDigit(digit) {
		return ErrDigit
	}

	g[row][column].digit = digit
	return nil
}

func (g *Grid) Reset(row, column int) error {

	return g.Set(row, column, empty)
}

func NewSudoku(s [rows][columns]int8) *Grid {
	var g Grid
	for i, column := range s {
		for j, digit := range column {
			g[i][j].digit = digit
			if digit != empty {
				g[i][j].fixed = true
			}
		}
	}
	return &g
}

func (g *Grid) initialDigit(row, column int) bool {
	return g[row][column].fixed
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit >= 0 && digit <= 9
}

func main() {

	s := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	err := s.Set(0, 2, 10)

	fmt.Println("Set")
	for i, row := range s {
		fmt.Println()
		for j, column := range row {
			fmt.Printf("%d ", column.digit)
			if i == 8 && j == 8 {
				fmt.Println()
			}
		}
	}
	/*
		err = s.Reset(8, 6)

		fmt.Println("Reset")
		for i, row := range s {
			fmt.Println()
			for j, column := range row {
				fmt.Printf("%d ", column.digit)
				if i == 8 && j == 8 {
					fmt.Println()
				}
			}
		}
	*/
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		os.Exit(1)
	}

}
