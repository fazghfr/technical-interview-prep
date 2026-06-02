package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	ans := -1

	for right >= left {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 91))
}
