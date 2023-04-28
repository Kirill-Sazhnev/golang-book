package main

import (
	"fmt"
	"strconv"
	"strings"
)

func MoveZeros(arr []int) []int { //5kyu
	var intArr []int
	var zeroArr []int

	for _, v := range arr {
		switch v != 0 {
		case true:
			intArr = append(intArr, v)
		case false:
			zeroArr = append(zeroArr, v)
		}
	}
	return append(intArr, zeroArr...)
}

func SpinWords(str string) string { //6kyu
	words := strings.Fields(str)
	sdrow := make([]string, len(words))
	for i, word := range words {
		if len(word) > 4 {
			drow := make([]byte, len(word))
			for j, letter := range word {
				drow[len(word)-j-1] = byte(letter)
			}
			sdrow[i] = string(drow)
		} else {
			sdrow[i] = word
		}
	}
	return strings.Join(sdrow, " ")
} // SpinWords

func ArrayDiff(a, b []int) []int { //6kyu

	for _, valB := range b {
		difA := make([]int, 0)
		for _, valA := range a {
			if valB != valA {
				difA = append(difA, valA)
			}
		}
		a = difA
	}
	return a
}

func FindNb(m int) int { //6kyu
	currNb, cubes := 0, 0

	for i := 1; currNb < m; i++ {
		currNb += i * i * i
		cubes++
	}

	if currNb == m {
		return cubes
	}
	return -1
}

func Revrot(s string, n int) string { // 6kyu
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
	fmt.Println(FindNb(100))
}

// ([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}) // returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }
