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

func main() {

	n := big.NewInt(13)
	m := big.NewInt(120)

	fmt.Println(Height(n, m))
	fmt.Println()
}
