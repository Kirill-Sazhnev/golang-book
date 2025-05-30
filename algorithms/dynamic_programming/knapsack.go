package main

func findMaxKnapsackProfitHelperV2(capacity int, weights []int, values []int, n int) int {
	// Base cases
	if n == 0 || capacity == 0 {
		return 0
	}

	// Recursive cases
	// If the weight of the nth item is less than capacity, then:
	currentWeight := weights[n-1]
	if currentWeight <= capacity {
		// We either include the item and deduct the weight of the item from the knapsack capacity (to get the remaining capacity)
		// or we don't include the item at all. We pick the option that yields the highest value.
		currentValue := values[n-1]
		remainingCapacity := capacity - currentWeight
		valueAfterCurrentItem := findMaxKnapsackProfitHelperV2(remainingCapacity, weights, values, n-1)
		valueSkippingCurrentItem := findMaxKnapsackProfitHelperV2(capacity, weights, values, n-1)

		return max(currentValue+valueAfterCurrentItem, valueSkippingCurrentItem)
	} else {
		// Item can't be added to our knapsack if its weight is greater than the capacity
		return findMaxKnapsackProfitHelperV2(capacity, weights, values, n-1)
	}
}

func findMaxKnapsackProfitV1(capacity int, weights []int, values []int) int {
	n := len(weights)
	return findMaxKnapsackProfitHelperV2(capacity, weights, values, n)
}

func maxV1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxKnapsackProfitHelperV1(capacity int, weights []int, values []int, n int, table [][]int) int {
	// Base case
	if n == 0 || capacity == 0 {
		return 0
	}

	// If we have already solved this subproblem, fetch the result from memory
	if table[n][capacity] != -1 {
		return table[n][capacity]
	}

	currentWeight := weights[n-1]
	// Otherwise, we solve it and save the result in our lookup table
	if currentWeight <= capacity {
		remainingCapacity := capacity - currentWeight
		currentValue := values[n-1]
		valueAfterCurrentItem := findMaxKnapsackProfitHelperV1(remainingCapacity, weights, values, n-1, table)
		valueSkippingCurrentItem := findMaxKnapsackProfitHelperV1(capacity, weights, values, n-1, table)

		table[n][capacity] = max(currentValue+valueAfterCurrentItem, valueSkippingCurrentItem)
		return table[n][capacity]
	}
	valueSkippingCurrentItem := findMaxKnapsackProfitHelperV1(capacity, weights, values, n-1, table)
	table[n][capacity] = valueSkippingCurrentItem
	return table[n][capacity]
}

func findMaxKnapsackProfitV2(capacity int, weights []int, values []int) int {
	n := len(weights)
	dp := make([]int, capacity+1)

	for i := 0; i < n; i++ {
		currWeight := weights[i]
		currValue := values[i]
		for currentCapacity := capacity; currentCapacity >= currWeight; currentCapacity-- {
			remainingCapacity := currentCapacity - currWeight
			maxForRemainingCapacity := dp[remainingCapacity]
			maxForCurrentCapacity := dp[currentCapacity]

			maxVal := max(currValue+maxForRemainingCapacity, maxForCurrentCapacity)
			dp[currentCapacity] = maxVal
		}
	}

	return dp[capacity]
}

func findMaxKnapsackProfit(capacity int, weights []int, values []int) int {
	data := make([][]int, len(weights)+1)
	for i := range data {
		data[i] = make([]int, capacity+1)
	}

	for i := 0; i < len(weights); i++ {
		weight := weights[i]
		value := values[i]
		for currentWeight := weight; currentWeight <= capacity; currentWeight++ {
			remainingCapacity := currentWeight - weight
			remainingCapacityVal := data[i][remainingCapacity]
			previousItemMaxValue := data[i][currentWeight]
			currValueAndValueOfRemainingCapacity := value + remainingCapacityVal
			data[i+1][currentWeight] = max(currValueAndValueOfRemainingCapacity, previousItemMaxValue)
		}
	}
	return data[len(weights)][capacity]
}
