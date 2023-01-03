package main

import (
	"fmt"
	"math"
)

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}

	return "Odd"
}

func PositiveSum(numbers []int) int {
	if numbers == nil {
		return 0
	}
	var sum int
	for _, v := range numbers {
		if v > 0 {
			sum += v
		}
	}
	return sum
}

func Solution(word string) string {
	str := ""
	for i := len(word); i > 0; i-- {
		str += string(word[i-1])
	}
	return str
}

func FindNextSquare(sq int64) int64 { //7kyu
	root := math.Sqrt(float64(sq))
	if root > float64(math.Floor(root)) {
		return -1
	}
	next := int64(math.Pow(root+1, 2))
	return next
}

func main() {
	fmt.Printf("%v", FindNextSquare(10))
}
