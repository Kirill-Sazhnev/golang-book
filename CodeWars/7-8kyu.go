package main

import (
	"math"
	"strconv"
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

func Quadratic(x1, x2 int) [3]int {
	var a, b, c int = 1, (-x2 - x1), (x1 * x2)
	return [3]int{a, b, c}
}

func SumOfIntegersInString(strng string) int { //7kyu
	var sum int

	for i := 0; i < len(strng); i++ {
		if strng[i] <= '9' && strng[i] >= '0' {
			numb, err := strconv.Atoi(string(strng[i]) + nextSymb(i+1, strng))
			if err != nil {
				panic("Could not convert string to int")
			}
			i = counter - 1
			sum += numb
		}
	}
	return sum
}

var counter int

func nextSymb(ind int, strSlice string) string { //part of SumOfIntegersInString

	if ind < len(strSlice) && strSlice[ind] >= '0' && strSlice[ind] <= '9' {
		return string(strSlice[ind]) + nextSymb(ind+1, strSlice)
	}
	counter = ind
	return ""
}

func SumOfIntegersInStr(strng string) (sum int) { //refactoring
	digit := 0
	for _, r := range strng {
		if r <= '9' && r >= '0' {
			digit = digit*10 + int(r-'0')
			continue
		}

		if digit > 0 {
			sum += digit
			digit = 0
		}
	}
	sum += digit
	return
}

func Cats(start, finish int) int {
	delta := finish - start
	jumps := 0.0
	if delta > 2 {
		jumps = math.Floor(float64(delta/3)) + float64(delta%3)
	} else {
		jumps = float64(delta)
	}
	return int(jumps)
}

/*
func main() {
	fmt.Println(Cats(1, 5))
}
*/
