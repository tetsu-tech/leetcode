package main

import (
	"fmt"
	"math"
)

func main() {
	result := coinChange([]int{1, 2, 5}, 11)
	fmt.Printf("result: %d\n", result)
}

// coins: array of coins
// target_amount: represents a total amount of money
func coinChange(coins []int, target_amount int) int {
	// initalize array, [-1, -1, -1....]
	// number of element equals target_amount
	// dp[i] represents the minimum number of coins to make up number i
	// if dp[i] is -1, this number cannot be made with given array of coins
	dp := make([]int, target_amount+1)
	// initialize dp[i] = -1, new value will be assigned to dp[i] if coins can be made up
	for i := range dp {
		dp[i] = -1
	}

	// calculate the number of coins with Dynamic Programing
	for subtarget_amount := range dp {
		// amount 0 can always be made with no coin
		if subtarget_amount == 0 {
			dp[0] = 0
			continue
		}

		for _, coin := range coins {
			if subtarget_amount-coin < 0 || dp[subtarget_amount-coin] == -1 {
				continue
			}

			if dp[subtarget_amount] == -1 {
				dp[subtarget_amount] = dp[subtarget_amount-coin] + 1
			} else {
				dp[subtarget_amount] = int(math.Min(float64(dp[subtarget_amount]), float64(dp[subtarget_amount-coin]+1)))
			}
		}
		fmt.Printf("===subtarget_amount=%d, dp[subtarget_amount]=%d\n", subtarget_amount, dp[subtarget_amount])
	}

	return dp[target_amount]
}
