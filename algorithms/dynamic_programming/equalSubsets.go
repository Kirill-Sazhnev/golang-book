package main

import "fmt"

// of the DP table becomes `true`.
func canPartitionArrayVerbose(nums []int) bool {
	//-------------------------------------------------------
	// 1. Compute the overall sum.
	//-------------------------------------------------------
	totalSum := 0
	for _, value := range nums {
		totalSum += value
	}
	fmt.Printf("Total sum = %d\n", totalSum)

	//-------------------------------------------------------
	// 2. If the sum is odd we can stop: two equal integers
	//    can never add up to an odd number.
	//-------------------------------------------------------
	if totalSum%2 == 1 {
		fmt.Println("Total is odd ⇒ impossible to partition")
		return false
	}
	targetSum := totalSum / 2
	fmt.Printf("Each subset must reach = %d\n\n", targetSum)

	//-------------------------------------------------------
	// 3. Dynamic-programming table:
	//    reachable[ s ] == true  ⇔  some subset hits sum s.
	//    We need only 0…targetSum (inclusive).
	//-------------------------------------------------------
	reachable := make([]bool, targetSum+1)
	reachable[0] = true // “empty subset” always makes 0

	//-------------------------------------------------------
	// 4. Process each number exactly once (0/1 knapsack).
	//-------------------------------------------------------
	for idx, number := range nums {
		fmt.Printf("#%d) considering number %d\n", idx+1, number)

		// Traverse *backwards* so we do not pick 'number' twice.
		for s := targetSum; s >= number; s-- {
			if !reachable[s] && reachable[s-number] {
				reachable[s] = true
				fmt.Printf("    reach sum %2d  by adding %d to previous %2d\n",
					s, number, s-number)
			}
		}

		// Optional: show a snapshot of the table after this round
		fmt.Printf("    reachable sums now true: ")
		for s, ok := range reachable {
			if ok {
				fmt.Printf("%d ", s)
			}
		}
		fmt.Println("\n")

		// Early exit — as soon as targetSum is reachable, we are done.
		if reachable[targetSum] {
			fmt.Println("Target reached! Partition exists.\n")
			return true
		}
	}

	//-------------------------------------------------------
	// 5. Finished the loop: answer depends on reachable[targetSum].
	//-------------------------------------------------------
	if reachable[targetSum] {
		fmt.Println("Partition exists.\n")
	} else {
		fmt.Println("Partition does NOT exist.\n")
	}
	return reachable[targetSum]
}

//func main() {
//	fmt.Println("Example 1:", canPartitionArrayVerbose([]int{1, 5, 11, 5}))
//	fmt.Println("Example 2:", canPartitionArrayVerbose([]int{1, 2, 3, 5}))
//}
