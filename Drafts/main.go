package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Revrot(s string, n int) string {
	if n < 1 {
		return ""
	}
	subStrgs := make([]string, 0)

	for i := 0; i <= len(s)-n; i += n {
		subStrgs = append(subStrgs, s[i:i+n])
	}

	for i, subStr := range subStrgs {
		if cubeSum(subStr)%2 == 0 {
			subStrgs[i] = reverse(subStr)
		} else {
			bytes := []byte(subStrgs[i])
			bytes = append(bytes[1:], bytes[0])
			subStrgs[i] = string(bytes)
		}
	}
	return strings.Join(subStrgs, "")
}

func cube(n int) int {
	return n * n * n
}

func cubeSum(nums string) int {
	sum := 0
	for _, val := range nums {
		num, _ := strconv.Atoi(string(val))
		sum += cube(num)
	}
	return sum
}

func reverse(s string) string {
	ln := len(s)
	str := []byte(s)
	for i := 0; i < ln/2; i++ {
		str[i], str[ln-i-1] = str[ln-i-1], str[i]
	}
	return string(str)
}

func main() {
	fmt.Println(Revrot("123456987653", 6), "234561356789")
}
