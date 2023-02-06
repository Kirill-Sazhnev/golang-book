package main

import "fmt"

func sumArr(data []int) int {
	total := 0

	for _, value := range data {
		total += value
	}
	//Implement your solution here

	return total
}

func SequentialSearch(data []int, value int) bool {
	//Implement your solution here
	for _, elemnt := range data {
		if elemnt == value {
			return true
		}
	}
	return false //Return false if value not found
}

func BinarySearch(data []int, value int) bool {
	//Implement your solution here
	slice := data
	for i := len(data) / 2; len(slice) > 0; i = len(slice) / 2 {

		switch {
		case value == slice[i]:
			return true
		case value > slice[i]:
			slice = slice[i+1:]
		case value < slice[i]:
			slice = slice[:i]
		}

	}
	fmt.Println(len(slice))
	return false //Return false if value not found
}

func MaxSubArraySum(data []int) int {
	sumSlice := []int{}
	maxVal := 0

	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			if sumArr(data[i:j+1]) > maxVal {
				maxVal = sumArr(data[i : j+1])
			}
		}
	}

	//Implement your solution here
	fmt.Println(sumSlice)
	return maxVal
}

func RotateArray(data []int, k int) {
	//Implement your solution here
	k = k % len(data)
	newArr := append(data[k:], data[:k]...)
	for i, val := range newArr {
		data[i] = val
	}
	fmt.Println(data)
	//You must make changes in the given array
}

func WaveArray(arr []int) {
	//Implement your solution here
	lenarr := len(arr)
	for i := 0; i < lenarr-1; i += 2 {
		if i+3 > lenarr {
			findMin(arr[i : i+2])
			break
		}
		findMin(arr[i : i+3])
	}
	//Update the same array
}

func findMin(arr []int) {
	minV, minI := arr[0], 0
	for i := 1; i < len(arr); i++ {
		if minV > arr[i] {
			minV, minI = arr[i], i
		}
	}
	temp := arr[1]
	arr[1] = minV
	arr[minI] = temp
}

func indexArray(arr []int, size int) {
	numbers := make(map[int]uint)
	for _, val := range arr {
		numbers[val]++
	}
	for i := 0; i < size; i++ {
		if _, ok := numbers[i]; ok {
			arr[i] = i
			numbers[i]--
		} else {
			arr[i] = -1
			numbers[-1]--
		}
	}

	//Note: Please make changes in the given array
}

func Sort1toN(arr []int, size int) {
	temp := 0
	for i := 0; i < size; i++ {
		for arr[i] != i+1 {
			temp = arr[i]
			arr[i] = arr[temp-1]
			arr[temp-1] = temp
		}
	}
}

func main() {
	arr := []int{3, 5, 2, 1, 4}
	Sort1toN(arr, 5)
	fmt.Println(arr)
}
