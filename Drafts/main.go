package main

import "fmt"

func Passanger(n, m int) {
	mid := m / 2
	if m%2 != 0 {
		mid++
	}
	next := mid
	for i := 1; i <= n; i++ {
		switch {
		case next == mid:
			fmt.Println(next)
			if m%2 == 0 {
				next++
			} else {
				next--
			}
		case i%m == 0:
			fmt.Println(next)
			next = mid

		case next > mid:
			fmt.Println(next)
			next = next - i%m

		case next < mid:
			fmt.Println(next)
			next = next + i%m

		}
	}
}

func Stall(n, k int) {
	max := (n-(k-1))*(n-(k-1)) + (k - 1)
	min := 0
	step := n / k
	Sleft := n
	left := n % k
	for Sleft > 0 {
		if left > 0 {
			Sleft -= step + left/left
			min += (step + left/left) * (step + left/left)
			left--
		} else {
			Sleft -= step
			min += step * step
		}

	}
	fmt.Printf("%v %v\n", min, max)
}

func main() {
	arr := []int{3, 5, 2, 1, 4}
	fmt.Println(arr)
}
