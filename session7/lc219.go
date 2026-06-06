package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}

	freq := make(map[int]int)

	left := 0
	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++

		if right-left > k {
			freq[nums[left]]--
			left++
		}

		if freq[nums[right]] > 1 {
			return true
		}
	}

	return false

}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(containsNearbyDuplicate(nums, 2))
}
