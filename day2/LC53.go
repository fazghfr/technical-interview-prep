package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	runningSum := 0
	globalSum := math.MaxInt32 * -1
	for _, v := range nums {
		if runningSum > 0 {
			// extend
			runningSum += v
		} else {
			// start fresh
			runningSum = v
		}
		globalSum = max(globalSum, runningSum)
	}
	return globalSum
}

func main() {
	// arr := []int{-2,1,-3,4,-1,2,1,-5,4}
	arr := []int{5, 4, -1, 7, 8}
	ans := maxSubArray(arr)
	fmt.Println(ans)
}
