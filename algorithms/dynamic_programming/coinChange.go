package main

func coinChangeV3(coins []int, total int) int {
	if total == 0 {
		return 0
	}
	l := len(coins)
	countersArray := make([]int, total+1)
	for _, coin := range coins {
		countersArray[coin] = 1
	}
	for i := range coins {
		findChangeV3(coins, l-(i+1), 0, total, 0, countersArray)
	}
	res := countersArray[total]
	if res == 0 {
		return -1
	}
	return res
}

func findChangeV3(coins []int, ix, sum, remainder, counter int, cm []int) {
	if remainder <= 0 {
		return
	}
	updateCm(cm, sum, counter)

	currentVal := coins[ix]
	if currentVal == remainder {
		updateCm(cm, sum+currentVal, counter+1)
		return
	}
	steps := cm[remainder]
	if steps != 0 {
		updateCm(cm, sum+remainder, counter+steps)
		return
	}
	l := len(coins)
	for i := range coins {
		findChangeV3(coins, l-(i+1), sum+currentVal, remainder-currentVal, counter+1, cm)
	}
}

func updateCm(cm []int, sum, counter int) {
	if counter == 0 {
		return
	}
	oldCount := cm[sum]
	if oldCount != 0 {
		cm[sum] = minCounter(counter, oldCount)
	} else {
		cm[sum] = counter
	}
}

func coinChangeV2(coins []int, total int) int {
	if total == 0 {
		return 0
	}
	minCnt := -1
	for i := range coins {
		findChangeV2(coins[i:], 0, total, 0, &minCnt)
	}
	return minCnt
}

func findChangeV2(coins []int, ix, remainder, counter int, minCtr *int) {
	if remainder <= 0 || exceededCounter(counter, *minCtr) {
		return
	}
	currentVal := coins[ix]
	if currentVal == remainder {
		*minCtr = minCounter(counter+1, *minCtr)
		return
	}
	for i := range coins {
		findChangeV2(coins, i, remainder-currentVal, counter+1, minCtr)
	}
}

func coinChange(coins []int, total int) int {
	cm := make(map[int]int, total)
	return findChange(coins, total, cm)
}

func findChange(coins []int, total int, counterMap map[int]int) int {
	if val, ok := counterMap[total]; ok {
		return val
	}
	if total == 0 {
		return 0
	}
	if total < 0 {
		return -1
	}

	minCnt := -1

	for _, coin := range coins {
		remainder := total - coin
		result := findChange(coins, remainder, counterMap)
		if result == -1 {
			continue
		}
		minCnt = minCounter(result+1, minCnt)
	}

	counterMap[total] = minCnt
	return minCnt
}

func coinChangeWithTab(coins []int, total int) int {
	cm := make([]int, total+1)

	for i := 0; i < total; i++ {
		currentCounter := cm[i]
		if currentCounter == 0 && i != 0 {
			continue
		}
		for _, coin := range coins {
			if i+coin > total {
				continue
			}
			nextCounter := cm[i+coin]
			if nextCounter == 0 || nextCounter > currentCounter+1 {
				cm[i+coin] = currentCounter + 1
			}
		}
	}
	if cm[total] == 0 {
		return -1
	}
	return cm[total]
}

func exceededCounter(a, b int) bool {
	if b == -1 || b > a {
		return false
	}
	return true
}

func minCounter(a, b int) int {
	if b == -1 || a < b {
		return a
	}
	return b
}
