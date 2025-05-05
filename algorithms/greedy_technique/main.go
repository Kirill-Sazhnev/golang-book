package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 2, 1, 0, 4}
	fmt.Println(jumpGame(arr))
}

func gasStationJourney(gas []int, cost []int) int {
	var (
		totalGas     int
		totalCost    int
		totalSubGas  int
		totalSubCost int
		start        int
		end          int
	)
	for i := 0; i <= len(gas)*2; i++ {
		if i == len(gas) && totalGas < totalCost {
			return -1
		}
		ix := i % len(gas)
		currentGas := gas[ix]
		currentCost := cost[ix]
		totalGas += currentGas
		totalCost += currentCost

		end++
		totalSubGas += currentGas
		totalSubCost += currentCost

		if totalSubGas >= totalSubCost {
			if (end - start) == len(gas) {
				return start
			}
			continue
		}

		for totalSubGas < totalSubCost {
			totalSubGas -= gas[start]
			totalSubCost -= cost[start]
			start++
		}
	}

	return -1
}

func gasStationJourneyV1(gas []int, cost []int) int {
	var (
		totalGas     int
		totalCost    int
		totalSubGas  int
		totalSubCost int
		start        int
		end          int
	)
	for i := range gas {
		currentGas := gas[i]
		currentCost := cost[i]
		totalGas += currentGas
		totalCost += currentCost

		end++
		totalSubGas += currentGas
		totalSubCost += currentCost

		if totalSubGas >= totalSubCost {
			if (end - start) == len(gas) {
				return start
			}
			continue
		}

		for totalSubGas < totalSubCost {
			totalSubGas -= gas[start]
			totalSubCost -= cost[start]
			start++
		}
	}
	if totalGas < totalCost {
		return -1
	}

	for i := range gas {
		currentGas := gas[i]
		currentCost := cost[i]

		end++
		totalSubGas += currentGas
		totalSubCost += currentCost

		if totalSubGas >= totalSubCost {
			if (end - start) == len(gas) {
				return start
			}
			continue
		}

		for totalSubGas < totalSubCost {
			totalSubGas -= gas[start]
			totalSubCost -= cost[start]
			start++
		}
	}
	return -1
}

func gasStationJourneyV2(gas []int, cost []int) int {
	for i := range gas {

		var (
			steps     int
			available int
		)
		for j := range cost {
			ix := (i + j) % len(gas)
			available += gas[ix]
			required := cost[ix]
			if available < required {
				break
			}
			available -= required
			steps++
		}
		if steps == len(cost) {
			return i
		}
	}
	return -1
}

func rescueBoats(people []int, limit int) int {

	sort.Ints(people)
	leftIx, rightIx := 0, len(people)-1
	var boats int
	for leftIx <= rightIx {
		left, right := people[leftIx], people[rightIx]
		boats++
		switch {
		case left+right <= limit:
			leftIx++
			rightIx--
		case left > right:
			leftIx++
		case left < right:
			rightIx--
		default:
			if leftIx != rightIx {
				boats++
			}
			leftIx++
			rightIx--
		}
	}
	return boats
}

func jumpGame(nums []int) bool {

	current := nums[0]
	ix := 0
	for ix < len(nums)-1 {
		if current == 0 {
			return false
		}
		steps := ix + current
		if steps >= len(nums) {
			return true
		}
		ix, current = findMax(nums, ix+1, steps)

	}
	return ix >= len(nums)-1
}

func findMax(nums []int, start, end int) (int, int) {
	maxValue := nums[start]
	ix := start
	for i := start; i <= end && i < len(nums); i++ {
		if nums[i] > maxValue {
			ix = i
			maxValue = nums[i]
		}
	}
	return ix, maxValue
}
