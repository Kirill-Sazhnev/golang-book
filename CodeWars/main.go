package main

import (
	"fmt"
	"math/big"
)

func main() {

	n := big.NewInt(13)
	m := big.NewInt(120)

	fmt.Println(Height(n, m))
	fmt.Println()
}
