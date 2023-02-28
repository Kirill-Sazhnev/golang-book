package main

import (
	"fmt"
)

func Factorial(i int) int {
	if i <= 1 {
		i = 1
	} else {
		i = i * Factorial(i-1)
	}
	return i
}

func printInt(number int) {

	if number > 15 {
		printInt(number / 16)
		printInt(number % 16)
	}
	switch number {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
		fmt.Print(number)
	case 10:
		fmt.Print("A")
	case 11:
		fmt.Print("B")
	case 12:
		fmt.Print("C")
	case 13:
		fmt.Print("D")
	case 14:
		fmt.Print("E")
	case 15:
		fmt.Print("F")
	}
}

func GCD(m int, n int) int {
	if m < n {
		m, n = n, m
	}

	if m%n == 0 {
		return n
	}

	return GCDrec(m, n, n-1)
}

func GCDrec(m, n, div int) int {

	if m%div == 0 && n%div == 0 {
		return div
	}
	return GCDrec(m, n, div-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func Permutation(data []int, i int, length int) {
	/*
		// base condition DO NOT ALTER IT
		if length == i {
			temp := "{"
			for k := 0; k < length; k++ {
				temp += strconv.Itoa(data[k])
				temp += " "
			}
			temp += "} "
			fmt.Print(temp)
			return
		}
	*/
	fmt.Println(data)

	j := (i + 1) % length

	data[i%length], data[j] = data[j], data[i%length]

	if i < length*(length-1)-1 {
		Permutation(data, i+1, length)
	}
	return

}

func towerOfHanoi(num int, src byte, dst byte, temp byte) {
	// write some code here

	if num < 1 {
		return
	}
	towerOfHanoi(num-1, src, temp, dst)
	fmt.Printf("Move disk %d from peg %c to peg %c \n", num, src, dst)
	towerOfHanoi(num-1, temp, dst, src)
}

func BinarySearchRecursive(data []int, value int) bool {
	size := len(data)
	return BinarySearchRecursiveUtil(data, 0, size-1, value)
}

func BinarySearchRecursiveUtil(data []int, low int, high int, value int) bool {

	//Write your code here
	midi := (high-low)/2 + low
	midv := data[midi]
	fmt.Println(midi, midv)

	switch {
	case value == midv:
		return true
	case value > midv && low < high:
		return BinarySearchRecursiveUtil(data, midi+1, high, value)
	case value < midv && low < high:
		return BinarySearchRecursiveUtil(data, low, midi, value)
	default:
		break
	}
	return false
}

func main() {}
