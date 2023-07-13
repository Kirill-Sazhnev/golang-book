package main

import (
	"fmt"
	"math/big"
)

func calc(n int) *big.Int {
	res := big.NewInt(1)
	one, two := big.NewInt(1), big.NewInt(2)
	for i := 1; i < n; i++ {
		res.Mul(res, two).Add(res, one)
	}
	return res
}

func Height(n, m *big.Int) *big.Int {
	nInt, mInt := n.Int64(), m.Int64()

	if nInt < 1 || mInt < 1 {
		return big.NewInt(0)
	}
	if nInt > mInt {
		return calc(int(mInt))
	}

	one, two := big.NewInt(1), big.NewInt(2)
	step := (int(mInt - nInt)) + 1

	arr1, arr2 := make([]big.Int, step), make([]big.Int, step)

	for i := range arr1 {
		arr1[i] = *big.NewInt(int64(i) + 1)
	}

	start := big.NewInt(1)
	for i := 1; i < int(nInt); i++ {

		start.Mul(start, two).Add(start, one)
		arr2[0] = *start

		for j := 1; j < step; j++ {
			arr2[j].Add(&arr1[j], &arr2[j-1]).Add(&arr2[j], one)
		}
		arr1, arr2 = arr2, arr1
	}
	return &arr1[step-1]
}

func Spiralize(size int) [][]int { // 3 kyu
	spiral := make([][]int, size)
	for i := range spiral {
		spiral[i] = make([]int, size)
	}

	field := make([][]int, size)
	for i := range field {
		field[i] = spiral[i]
	}

start:
	if len(field) > 1 {
		for i := 0; i < len(field[0]); i++ {
			field[0][i] = 1
			field[1][i] = 0
		}
		field[1][len(field[1])-1] = 1
		field = field[2:]
	}

	for i := 0; i < len(field) && len(field[0]) > 1; i++ {
		field[i][len(field[i])-1] = 1
		if i == len(field)-1 {
			field[i][len(field[i])-2] = 1
		} else {
			field[i][len(field[i])-2] = 0
		}
		field[i] = field[i][:len(field[i])-2]
	}

	if len(field) > 1 {
		for i := 0; i < len(field[len(field)-1]); i++ {
			if i == 0 {
				field[len(field)-2][i] = 1
				field[len(field)-1][i] = 1
			} else {
				field[len(field)-2][i] = 0
				field[len(field)-1][i] = 1
			}
		}
		field = field[:len(field)-2]
	}

	for i := 0; i < len(field) && len(field[0]) > 1; i++ {
		if i == 0 {
			field[i][0] = 1
			field[i][1] = 1
		} else {
			field[i][0] = 1
			field[i][1] = 0
		}
		field[i] = field[i][2:]
	}

	ln := len(field)
	switch ln {
	case 0:
		return spiral
	case 1:
		field[0][0] = 1
		return spiral
	default:
		goto start
	}
}

func main() {

	n := big.NewInt(13)
	m := big.NewInt(120)

	fmt.Println(Height(n, m))
	fmt.Println()
}
