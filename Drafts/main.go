package main

import (
	"fmt"
)

func PrimeFactors(n int) (res []int) { // 6 kyu
	var factors []int
	rem := n
	for i := 2; i <= rem/i; i++ {
		if rem%i == 0 {
			factors = append(factors, i)
			rem = rem / i
			i = 1
		}
	}
	if rem > 1 {
		factors = append(factors, rem)
	}

	return factors
}

func main() {

	fmt.Println(PrimeFactors(12))
}
