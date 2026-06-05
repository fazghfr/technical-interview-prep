package main

import "fmt"

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
// You must write an algorithm that runs in O(n) time.

// Example 1:

// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
// Example 2:

// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9
// Example 3:

// Input: nums = [1,0,1,2]
// Output: 3

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	starts := make(map[int]bool)

	for _, v := range nums {
		m[v] = true
	}

	for _, v := range nums {
		if !m[v-1] {
			starts[v] = true
		}
	}

	ans := 0
	for k, _ := range starts {
		curr := 0
		for m[k] {
			curr++
			k++
		}
		ans = max(ans, curr)
	}

	return ans
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
