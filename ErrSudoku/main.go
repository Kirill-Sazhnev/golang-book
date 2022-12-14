package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const rows, columns = 9, 9

var (
	ErrBounds  = errors.New("out of grid boundaries")
	ErrDigit   = errors.New("invalid digit")
	ErrInitial = errors.New("Initial digit cannot be changed")
)

type Grid [rows][columns]int8

type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error()) // Конвертирует ошибки в строки
	}
	return strings.Join(s, ", ")
}

func (g *Grid) Set(row, column int, digit int8, initial Grid) error {
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	} else if initialDigit(row, column, initial) {
		errs = append(errs, ErrInitial)
	}

	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}

	if len(errs) > 0 {
		return errs
	}

	g[row][column] = digit
	return nil
}

func (g *Grid) Reset(row, column int, initial Grid) error {

	return g.Set(row, column, 0, initial)
}

func NewSudoku(g [rows][columns]int8) Grid {
	return g
}

func initialDigit(row, column int, i Grid) bool {

	if i[row][column] != 0 {
		return true
	}
	return false
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
	var g Grid

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

	g = s

	err := g.Set(8, 6, 9, s)

	fmt.Println("Set")
	for _, row := range g {
		fmt.Println(row)
	}

	err = g.Reset(-8, 6, s)

	fmt.Println("Reset")
	for _, row := range g {
		fmt.Println(row)
	}

	if err != nil {
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
		os.Exit(1)
	}

}
