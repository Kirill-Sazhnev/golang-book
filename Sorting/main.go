package main

import "fmt"

func Partition01(arr []int, size int) int {
	//Write your code here
	start := 0
	end := size - 1
	count := 0 //Change the count variable as per needs

	for i := 0; i < size && start < end; i++ {
		switch {
		case arr[start] > arr[end]:
			arr[start], arr[end] = arr[end], arr[start]
			count++
		case arr[end] == 0:
			start++
		case arr[start] == 1:
			end--
		default:
			start++
			end--
		}
	}
	//Note: Update the same given array
	//Write your code here
	return count
}

func Partition012(arr []int, size int) {
	//Implement your solution here
	start := 0
	end := start + 1

	for i := 0; i < size && end < size; i++ {
		switch {
		case arr[start] < 1:
			start++
			end++
		case arr[end] < 1:
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end++
		default:
			end++
		}

	}
	bckStrt := size - 1
	bckNd := bckStrt - 1

	for i := 0; i < size && bckNd >= start; i++ {
		switch {
		case arr[bckStrt] > 1:
			bckStrt--
			bckNd--
		case arr[bckNd] > 1:
			arr[bckStrt], arr[bckNd] = arr[bckNd], arr[bckStrt]
			bckStrt--
			bckNd--
		default:
			bckNd--
		}
		//Note: Update the same given array
	}
}

func main() {
	arr := []int{0, 0, 2, 0, 2, 1, 0, 1}
	size := len(arr)
	Partition012(arr, size)
	fmt.Printf("%v", arr)
}
