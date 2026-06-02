package main

import (
	"fmt"
)

// similar to coin change
// amount -> totalsum/2 since if we have sum of any subsets equal to this, the rest will be automatically sum/2
// KNAPSACK
func canPartition(nums []int) bool {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}
	targetSum := totalSum / 2
	dp := make([]bool, targetSum+1)
	dp[0] = true
	for _, v := range nums {
		for i := targetSum; i >= v; i-- {
			dp[i] = dp[i] || dp[i-v]
		}
	}
	return dp[targetSum]
}

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}
