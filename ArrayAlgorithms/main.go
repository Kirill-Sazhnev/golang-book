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

func SmallestPositiveMissingNumber(arr []int, size int) int {

	min := -1
	j := 1
	for i := 0; i < size-1; i++ {
		fmt.Println(i, i+j, arr)

		switch {
		case i+j >= size:
			j = 1
		case arr[i] > arr[i+j]:
			arr[i], arr[i+j] = arr[i+j], arr[i]
			j++
			i--
		case arr[i] < arr[i+j]:
			j++
			i--
		}
	}
	for i, val := range arr {
		if val != i+1 {
			min = i + 1
			break
		}
	}
	return min //Return -1 if missing number not found
}

func MaxMinArr(arr []int, size int) {
	var maxArr []int
	maxArr = append(maxArr, arr[size/2:]...)
	var minArr []int
	minArr = append(minArr, arr[:size/2]...)

	maxi := len(maxArr) - 1
	mini := 0

	for i := 0; i < len(arr)-1; i += 2 {
		arr[i] = maxArr[maxi]
		maxi--
		arr[i+1] = minArr[mini]
		mini++
	}

	if size%2 != 0 {
		arr[size-1] = maxArr[0]
	}

}

func ArrayIndexMaxDiff(arr []int, size int) int {
	dist := 0
	for i := 0; i < size-1; i++ {

		for j := size - 1; j > i && (j-i) > dist; j-- {
			if arr[i] < arr[j] && (j-i) > dist {
				dist = j - i
				break
			}
		}
	}
	return dist
}

func ArrayIndexMaxDiff2(arr []int, size int) int {
	leftMin := make([]int, size)
	rightMax := make([]int, size)
	leftMin[0] = arr[0]
	i, j := 0, 0
	var maxDiff = 0
	for i = 1; i < size; i++ {
		if leftMin[i-1] < arr[i] {
			leftMin[i] = leftMin[i-1]
		} else {
			leftMin[i] = arr[i]
		}
	}
	rightMax[size-1] = arr[size-1]
	for i = size - 2; i >= 0; i-- {
		if rightMax[i+1] > arr[i] {
			rightMax[i] = rightMax[i+1]
		} else {
			rightMax[i] = arr[i]
		}
	}
	i = 0
	j = 0
	maxDiff = -1
	for j < size && i < size {
		if leftMin[i] < rightMax[j] {
			if maxDiff < j-i {
				maxDiff = j - i
			}
			j = j + 1
		} else {
			i = i + 1
		}
	}
	fmt.Println(leftMin, rightMax)
	return maxDiff
}

func maxPathSum(arr1 []int, size1 int, arr2 []int, size2 int) int {
	minArr := arr1
	maxArr := arr2
	maxSum := 0
	if sum(arr1) > sum(arr2) {
		minArr = arr2
		maxArr = arr1
	}

	for i := len(minArr) - 1; i >= 0; i-- {
		for j := 0; j < len(maxArr); j++ {
			if minArr[i] == maxArr[j] {
				if newSum := sum(minArr[:i]) + sum(maxArr[j:]); newSum > maxSum {
					maxSum = newSum
				}
			}
		}
	}

	return maxSum
}

func sum(arr []int) (sum int) {
	for _, val := range arr {
		sum += val
	}
	return sum
}

func Factorial(i int) int {
	if i == 2 {
		i = 2 * 1
	} else {
		i = i * Factorial(i-1)
	}
	return i
}

func main() {

	fmt.Println(Factorial(5))

}
