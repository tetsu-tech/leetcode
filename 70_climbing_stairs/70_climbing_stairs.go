package main

import (
	"fmt"
)

func main() {
	n := 10

	fmt.Printf("%d\n", numberSteps(n))
}

func numberSteps(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		fmt.Printf("i: %d , dp[i]: %d, dp[i-1]: %d, dp[i-2]: %d\n", i, dp[i], dp[i-1], dp[i-2])
	}

	return dp[n]
}
