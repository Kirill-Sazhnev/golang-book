package main

import (
	"fmt"
)

func main() {
	arr := [][]int{{0, 1, 0, 1}, {1, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}
	res := updateMatrix(arr)
	for _, row := range res {
		fmt.Println(row)
	}
}

func updateMatrix(mat [][]int) [][]int {
	res := make([][]int, len(mat))
	for i := range res {
		res[i] = make([]int, len(mat[0]))
	}
	for i := range res {
		for j := range res[i] {
			current := mat[i][j]
			if current == 0 {
				res[i][j] = current
				putOnesAround(res, i, j)
				continue
			}
			if res[i][j] == 1 {
				continue
			}
			res[i][j] = getMinBehind(res, i, j)
		}
	}

	for i := len(res) - 1; i >= 0; i-- {
		for j := len(res[0]) - 1; j >= 0; j-- {
			current := res[i][j]
			if current == 0 || current == 1 {
				continue
			}
			res[i][j] = getMinAround(res, i, j)
		}
	}
	return res
}

func getMinBehind(mat [][]int, x, y int) int {
	m := len(mat) + len(mat[0])
	up, left := m, m
	if x-1 >= 0 {
		up = mat[x-1][y] + 1
	}
	if y-1 >= 0 {
		left = mat[x][y-1] + 1
	}
	return min(up, left)
}

func getMinAround(mat [][]int, x, y int) int {
	m := len(mat) + len(mat[0]) - 1
	down, right := m, m
	if x+1 < len(mat) {
		down = mat[x+1][y] + 1
	}
	if y+1 < len(mat[0]) {
		right = mat[x][y+1] + 1
	}
	if down == m && right == m {
		return mat[x][y]
	}
	return min(down, right)
}

func putOnesAround(mat [][]int, x, y int) [][]int {
	if x+1 < len(mat) {
		mat[x+1][y] = 1
	}
	if x-1 >= 0 && mat[x-1][y] != 0 {
		mat[x-1][y] = 1
	}
	if y+1 < len(mat[0]) {
		mat[x][y+1] = 1
	}
	if y-1 >= 0 && mat[x][y-1] != 0 {
		mat[x][y-1] = 1
	}
	return mat
}

func updateMatrixRecursively(mat [][]int) [][]int {
	res := make([][]int, len(mat))
	for i := range res {
		res[i] = make([]int, len(mat[0]))
		for j := range res[i] {
			res[i][j] = -1
		}
	}

	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				res = traverseFromZeros(res, i, j, 0)
			}
		}
	}
	return res
}

func traverseFromZeros(res [][]int, x, y, previous int) [][]int {
	if x == len(res) || y == len(res[0]) || x < 0 || y < 0 {
		return res
	}
	current := res[x][y]
	switch {
	case current == -1, current > previous:
		res[x][y] = previous
	case current <= previous:
		return res
	}
	newCounter := res[x][y] + 1
	res = traverseFromZeros(res, x-1, y, newCounter)
	res = traverseFromZeros(res, x+1, y, newCounter)
	res = traverseFromZeros(res, x, y+1, newCounter)
	res = traverseFromZeros(res, x, y-1, newCounter)
	return res
}

func climbStairs(n int) int {
	counters := make([]int, n+1)
	counters[1] = 1
	for i := 2; i <= n; i++ {

		for j := 1; j <= 2; j++ {
			prevCounter := counters[i-j]
			counters[i] += prevCounter + 1
		}
	}
	return counters[n]
}

func countingBits(n int) []int {
	result := make([]int, n+1)
	pointer := 2
	delta := 1
	for i := 1; i <= n; i++ {
		switch {
		case i == pointer:
			result[i] = 1
			pointer *= 2
			delta *= 2
		case i%2 == 0:
			result[i] = result[i-delta] + 1
		default:
			result[i] = result[i-delta] + 1
		}
	}
	return result
}

func arrayGrow(a []int) bool {
	if len(a) == 5 {
		return true
	}
	a = append(a, 1)
	if arrayGrow(a) {
		return true
	}
	return false
}

func canPartitionArrayV3(nums []int) bool {
	sum := calculateSum(nums)
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	seenMap := make(map[int]bool, len(nums))
	return calculateArray(target, nums, seenMap)
}

func calculateArray(target int, array []int, seen map[int]bool) bool {
	l := len(array)
	for i := 0; i < l; i++ {
		current := array[i]
		remaining := target - current
		if remaining < 0 || seen[remaining] {
			continue
		}
		if remaining == 0 {
			return true
		}
		seen[remaining] = true

		newNums := make([]int, len(array[:i]), l-1)
		copy(newNums, array[:i])
		newNums = append(newNums, array[i+1:]...)

		if calculateArray(remaining, newNums, seen) {
			return true
		}
	}
	return false
}

func canPartitionArrayV1(nums []int) bool {
	sumMap := make(map[int]int, len(nums))
	return calculateArrayV1(0, nums, sumMap)
}

func calculateArrayV1(leftSum int, rightArray []int, sm map[int]int) bool {
	l := len(rightArray)
	for i := 0; i < l-1; i++ {
		left := rightArray[i]
		newSum := leftSum + left
		rightSum, ok := sm[newSum]
		if ok {
			continue
		}
		newNums := make([]int, len(rightArray[:i]), l-1)
		copy(newNums, rightArray[:i])
		newNums = append(newNums, rightArray[i+1:]...)
		rightSum = calculateSum(newNums)

		sm[newSum] = rightSum
		sm[rightSum] = newSum
		if newSum == rightSum {
			return true
		}
		if newSum > rightSum {
			continue
		}
		if calculateArrayV1(newSum, newNums, sm) {
			return true
		}
	}
	return false
}

func canPartitionArrayV2(nums []int) bool {
	sum := calculateSum(nums)
	sumArr := make([]int, sum)
	if sum%2 != 0 {
		return false
	}
	return calculateArrayV2(0, nums, sumArr)
}

func calculateArrayV2(leftSum int, rightArray []int, sm []int) bool {
	l := len(rightArray)
	for i := 0; i < l-1; i++ {
		left := rightArray[i]
		newSum := leftSum + left
		rightSum := sm[newSum]
		if rightSum != 0 {
			continue
		}
		newNums := make([]int, len(rightArray[:i]), l-1)
		copy(newNums, rightArray[:i])
		newNums = append(newNums, rightArray[i+1:]...)
		rightSum = calculateSum(newNums)

		sm[newSum] = rightSum
		sm[rightSum] = newSum
		if newSum == rightSum {
			return true
		}
		if newSum > rightSum {
			continue
		}
		if calculateArrayV2(newSum, newNums, sm) {
			return true
		}
	}
	return false
}

func calculateSum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func updateMap(countersMap map[int]int, sum, counter int) {
	if counter == 0 {
		return
	}
	oldCount, ok := countersMap[sum]
	if ok {
		countersMap[sum] = minCounter(counter, oldCount)
	} else {
		countersMap[sum] = counter
	}
}

func tribonacci(n int) int {
	triArr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i < 3 {
			triArr[i] = 1
			continue
		}
		triArr[i] = triArr[i-1] + triArr[i-2] + triArr[i-3]
	}
	return triArr[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxKnapsackProfitV0(capacity int, weights []int, values []int) int {
	if capacity == 0 || len(weights) == 0 || len(values) == 0 {
		return 0
	}
	maxVal := 0
	valMap := make(map[int]int)

	for i, weight := range weights {
		newCapacity := capacity - weight
		if newCapacity < 0 {
			continue
		}

		nextVal, ok := valMap[newCapacity]
		if !ok {
			nextVal = nextKnapsackProfit(newCapacity, i+1, weights, values, valMap)
		}
		currVal := values[i] + nextVal
		if currVal > maxVal {
			maxVal = currVal
		}
	}
	return maxVal
}

func nextKnapsackProfit(capacity, ix int, weights []int, values []int, valMap map[int]int) int {
	if capacity == 0 || len(weights) == ix || len(values) == ix {
		return 0
	}
	maxVal := 0

	for i := ix; i < len(weights); i++ {
		weight := weights[i]
		newCapacity := capacity - weight
		if newCapacity < 0 {
			continue
		}

		nextVal, ok := valMap[newCapacity]
		if !ok {
			nextVal = nextKnapsackProfit(newCapacity, i+1, weights, values, valMap)
		}
		currVal := values[i] + nextVal
		if currVal > maxVal {
			maxVal = currVal
		}
	}
	valMap[capacity] = maxVal

	return maxVal
}
