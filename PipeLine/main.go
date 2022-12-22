package main

import (
	"fmt"
	"strings"
)

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	letter := "London is a capital capital of the the Great Britian"
	go writer(c0, letter)
	go stringFilter(c0, c1)
	printString(c1)

}

func writer(cIn chan string, let string) {
	words := strings.Fields(let)
	for _, word := range words {
		cIn <- word
	}
	close(cIn)
}

func stringFilter(cIn, cOut chan string) {
	current := " "

	for nextValue := range cIn {

		if nextValue != current {
			cOut <- nextValue
			current = nextValue
		}

	}
	close(cOut)
}

func printString(cIn chan string) {
	for v := range cIn {
		fmt.Printf("%s ", v)
	}
}
