package main

import (
	"fmt"
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

func main() {
	fmt.Println(SpinWords("Hey fellow warriors"))
}

// ([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}) // returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }
