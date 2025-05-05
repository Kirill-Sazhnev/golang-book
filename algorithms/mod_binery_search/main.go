package main

import (
	"fmt"
)

func main() {
	//arr := []int{1, 2, 3, 4, 5, 6, 7}
	//arr := []int{4}
	//arr := []int{-10, -6, -4, -3}
	arr := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 8, 8}

	//arr := []int{-10000, -9664, -9462, -8703, -8642, -8169, -8138, -7955, -7795, -7676, -7657, -7473, -7422, -7325, -6861, -6803, -6503, -6409, -6408, -6336, -5830, -5336, -5059, -5002, -4988, -4981, -4889, -4731, -4429, -4319, -3896, -3823, -3660, -3336, -2528, -2136, -1747, -1700, -1651, -1643, -1506, -951, -676, -642, -590, -452, -172, 138, 226, 302, 401, 510, 552, 862, 907, 910, 915, 954, 1070, 1131, 1194, 1914, 2440, 2692, 2701, 2774, 3262, 3266, 3302, 3376, 3554, 3587, 3848, 4387, 4429, 4452, 4570, 4850, 5177, 5185, 5462, 5602, 5773, 5946, 6428, 6822, 6836, 6874, 6972, 7054, 7262, 7400, 7661, 8038, 8146, 8368, 8378, 8638, 9019, 9037, 9123, 9535, 9835, 10000}
	fmt.Println(singleNonDuplicate(arr))
}

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) <= 2 {
		return -1
	}
	mid := len(nums) / 2
	lMid, rMid := mid-1, mid+1
	if nums[mid] != nums[lMid] && nums[mid] != nums[rMid] {
		return nums[mid]
	}
	var (
		leftArr, rightArr []int
	)

	if nums[lMid] == nums[mid] {
		leftArr, rightArr = nums[:lMid], nums[rMid:]
	} else {
		leftArr, rightArr = nums[:mid], nums[mid:]
	}

	left, right := -1, -1
	if len(leftArr)%2 == 1 {
		left = singleNonDuplicate(leftArr)
	} else {
		right = singleNonDuplicate(rightArr)
	}

	// Replace this placeholder return statement with your code
	return max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findClosestElements(nums []int, k int, target int) []int {

	if len(nums) <= k {
		return nums
	}
	if target < nums[0] {
		return nums[0:k]
	}
	if target > nums[len(nums)-1] {
		return nums[len(nums)-k:]
	}

	ix := binarySearch(nums, target)
	left, right := ix-1, ix

	for right-left-1 < k {
		if left < 0 {
			right++
			continue
		}
		if right >= len(nums) {
			left--
			continue
		}

		lVal := abs(nums[left] - target)
		rVal := abs(nums[right] - target)
		if lVal <= rVal {
			left--
		} else {
			right++
		}
	}
	return nums[left+1 : right]
}

func findClosestElementsV2(nums []int, k int, target int) []int {
	if len(nums) == k {
		return nums
	}

	if target <= nums[0] {
		return nums[0:k]
	}

	if target >= nums[len(nums)-1] {
		return nums[len(nums)-k : len(nums)]
	}

	firstClosest := binarySearch(nums, target)

	windowLeft := firstClosest - 1
	windowRight := windowLeft + 1

	for (windowRight - windowLeft - 1) < k {
		if windowLeft == -1 {
			windowRight += 1
			continue
		}

		if windowRight == len(nums) || abs(nums[windowLeft]-target) <= abs(nums[windowRight]-target) {
			windowLeft -= 1
		} else {
			windowRight += 1
		}
	}

	return nums[windowLeft+1 : windowRight]
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func binarySearchRotated(nums []int, target int) int {

	start := 0
	end := len(nums)
	mid := (start + end) / 2
	pivot := len(nums)
	subNums := nums
	if isRotated(nums) {
		for {
			if nums[mid] == target {
				return mid
			}

			switch {
			case isRotated(nums[start:mid]):
				end = mid + 1
			case isRotated(nums[mid:end]):
				start = mid
			default:
				start, end = mid-1, mid+1
			}

			if len(nums[start:end]) == 2 {
				pivot = end - 1
				break
			}
			mid = (start + end) / 2
		}
	}

	delta := 0
	if target >= nums[0] {
		subNums = nums[:pivot]
	} else {
		subNums = nums[pivot:]
		delta = pivot
	}
	start = 0
	end = len(subNums)
	for {
		if len(subNums[start:end]) == 0 {
			return -1
		}
		mid = (start + end) / 2
		if subNums[mid] == target {
			return (mid + delta) % len(nums)
		}
		if target < subNums[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}
}

func isRotated(nums []int) bool {
	return nums[0] > nums[len(nums)-1]
}

func binarySearch(array []int, target int) int {
	// Initialize the left and right pointer
	left := 0
	right := len(array) - 1

	// Find the first closest element to target while left is less than or equal to right
	for left <= right {
		// Compute the middle index
		mid := (left + right) / 2

		// If the value at mid is equal to the target, return mid
		if array[mid] == target {
			return mid
		}

		// If the value at mid is less than target, move left forward
		if array[mid] < target {
			left = mid + 1
		} else { // If the value at mid is greater than target, move right backward
			right = mid - 1
		}
	}

	// If the target is not found, return the left index (position where the target should be inserted)
	return left
}
