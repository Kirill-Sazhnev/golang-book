package main

import "fmt"

func insertionSort(sl []int) []int {

	for i := 1; i < len(sl); i++ {
		tempV := sl[i]
		tempI := i
		for j := i - 1; j >= 0 && tempV < sl[j]; j-- {
			sl[tempI] = sl[j]
			tempI--
		}
		if tempI != i {
			sl[tempI] = tempV
		}
	}
	return sl
}

func bubbleSort(sl []int) []int {

	for i := 1; i < len(sl); i++ {
		for j := i; j > 0 && sl[j-1] > sl[j]; j-- {
			sl[j], sl[j-1] = sl[j-1], sl[j]
		}
	}
	return sl
}

func main() {
	sl := []int{4, 5, 2, 1, 3, 7, 6}
	fmt.Println(bubbleSort(sl))
}
