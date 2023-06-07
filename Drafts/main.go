package main

import (
	"fmt"
)

func Solution(sl []int) int { // 4 kyu

	ar := make([]int, len(sl))
	copy(ar, sl)
	cntr := 0
	i, j := len(ar)-2, len(ar)-1

	for cntr != len(ar)-1 {
		switch {
		case ar[i] > ar[j]:
			if ar[i]%ar[j] == 0 {
				ar[i] = ar[j]
			} else {
				ar[i] %= ar[j]
			}
			cntr = 0
		case ar[j] > ar[i]:
			if ar[j]%ar[i] == 0 {
				ar[j] = ar[i]
			} else {
				ar[j] %= ar[i]
			}
			cntr = 0
		default:
			cntr++
			if i == 0 {
				i = len(ar) - 1
			}
			i--
			j = i + 1
		}
	}
	return ar[0] * len(ar)
}

func main() {
	arr := []int{6, 9, 21}
	fmt.Println(Solution(arr))
}
