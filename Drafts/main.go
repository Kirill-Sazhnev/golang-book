package main

import (
	"fmt"
	"math/big"
)

var base = make([][]big.Int, 0)
var start *big.Int = big.NewInt(0)

func HeightB(n, m *big.Int) *big.Int {
	nInt := int(n.Int64())
	mInt := int(m.Int64())

	if nInt < 1 || mInt < 1 {
		result := big.NewInt(0)
		return result.Mul(m, n)
	}

	if nInt > mInt {
		return calc(int(mInt))
	}

	one := big.NewInt(1)
	two := big.NewInt(2)
	step := (mInt - nInt) + 1

	if len(base) == 0 {
		nTry := make([]big.Int, step)
		for i := range nTry {
			nTry[i] = *big.NewInt(int64(i) + 1)
		}
		base = append(base, nTry)
	}

	if len(base[0]) < step {
		for i := len(base[0]) + 1; i <= step; i++ {
			base[0] = append(base[0], *big.NewInt(int64(i)))
		}
	}
	ln := 0
	for i := 1; i < nInt; i++ {

		if len(base) == i {
			row := make([]big.Int, step+i)
			base = append(base, row)

			start.Mul(start, two).Add(start, one)
			base[i][i-1] = *start
			ln = 0
		} else {
			ln = len(base[i]) - i
		}

		for j := i + ln; j < step+i; j++ {
			if ln == 0 {
				base[i][j].Add(&base[i][j-1], &base[i-1][j-1]).Add(&base[i][j], one)

			} else {
				tmp := big.NewInt(0)
				tmp.Add(&base[i][j-1], &base[i-1][j-1]).Add(tmp, one)
				base[i] = append(base[i], *tmp)
			}
		}
	}
	return &base[nInt-1][mInt-1]
}

func calc(n int) *big.Int {
	res := big.NewInt(1)
	one := big.NewInt(1)
	two := big.NewInt(2)
	for i := 1; i < n; i++ {
		res.Mul(res, two).Add(res, one)
	}
	return res
}

func Height(n, m *big.Int) *big.Int {
	nInt := n.Int64()
	mInt := m.Int64()
	fmt.Println(nInt, mInt)

	if nInt < 1 || mInt < 1 {
		return big.NewInt(0)
	}
	if nInt > mInt {
		return calc(int(mInt))
	}

	one := big.NewInt(1)
	two := big.NewInt(2)
	step := (int(mInt - nInt)) + 1

	mEggs := make([][]big.Int, nInt)
	nTry := make([]big.Int, step)

	for i := range nTry {
		nTry[i] = *big.NewInt(int64(i) + 1)
	}
	mEggs[0] = nTry

	start := big.NewInt(0)
	for i := 1; i < int(nInt); i++ {

		mEggs[i] = make([]big.Int, step+i)
		start.Mul(start, two).Add(start, one)
		mEggs[i][i-1] = *start

		for j := i; j < step+i; j++ {
			mEggs[i][j].Add(&mEggs[i][j-1], &mEggs[i-1][j-1]).Add(&mEggs[i][j], one)
		}
	}
	return &mEggs[nInt-1][mInt-1]
}

func HeightInt(n, m int) int {

	if n < 1 {
		return m * n
	}
	res := 0
	for i := 1; i <= m; i++ {
		res += HeightInt(n-1, i-1)
	}
	return res + m
}

func Height2D(n, m int) int {
	mEggs := make([][]int, n)
	nTry := make([]int, m)
	for i := range nTry {
		nTry[i] = i + 1
	}
	mEggs[0] = nTry
	for i := 1; i < n; i++ {
		mEggs[i] = make([]int, m)
		mEggs[i][0] = 1
		for j := 1; j < m; j++ {
			mEggs[i][j] = mEggs[i][j-1] + mEggs[i-1][j-1] + 1
		}
	}
	return mEggs[n-1][m-1]
}

func main() {
	/*
		n := big.NewInt(19630)
		m := big.NewInt(19630)

		fmt.Println(HeightB(n, m))
		fmt.Println()

		n = big.NewInt(19790)
		m = big.NewInt(19900)

		fmt.Println(HeightB(n, m))

				for i := 1; i <= 10; i++ {
					for j := 1; j < 11; j++ {
						fmt.Printf("%-3d ", HeightInt(i, j))
					}
					fmt.Println()
				}
				fmt.Println()
	*/
	for i := 1; i <= 20; i++ {
		for j := 1; j < 13; j++ {
			fmt.Printf("%-4d ", Height2D(i, j))
		}
		fmt.Println()
	}
	//fmt.Println(HeightInt(6, 5))

}
