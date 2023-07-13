package main

import (
	"fmt"
	"strconv"
)

func NextBigger(n int) int { // 4 kyu

	nStr := strconv.Itoa(n)
	var res string
	numbers := make([]int, 0)

	for i := len(nStr) - 1; i > 0 && len(nStr) > 1; i-- {
		left, _ := strconv.Atoi(string(nStr[i-1]))
		right, _ := strconv.Atoi(string(nStr[i]))
		numbers = append(numbers, right)

		if left < right {
			res = nStr[:i-1] + nextNum(numbers, left)
			resN, _ := strconv.Atoi(res)
			return resN
		}

	}
	return -1
}

func nextNum(arr []int, num int) string {
	var numStr string
	for _, val := range arr {
		if val > num {
			numStr = strconv.Itoa(val) + numStr + strconv.Itoa(num)
			num = arr[(len(arr) - 1)]
			continue
		}
		numStr += strconv.Itoa(val)
	}
	return numStr
}

func main() {

	fmt.Println(NextBigger(144))
}
