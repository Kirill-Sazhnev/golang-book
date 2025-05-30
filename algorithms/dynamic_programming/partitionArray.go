package main

func canPartitionArray(nums []int) bool {
	sum := calculateSum(nums)
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	sumArray := make([][]bool, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		sumArray[i] = make([]bool, sum+1)
		sumArray[i][0] = true
	}

	for i, number := range nums {
		for growingSum := 0; growingSum <= target; growingSum++ {
			if sumArray[i][growingSum] {
				// copy previous result
				sumArray[i+1][growingSum] = true
				// add a new reachable result
				nextSum := growingSum + number
				sumArray[i+1][nextSum] = true
			}
		}
	}
	return sumArray[len(nums)][target]
}
