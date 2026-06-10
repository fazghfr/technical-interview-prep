package main

import "fmt"

func nextGreaterElements2(nums []int) []int {
	var stack []int
	ans := make([]int, len(nums))
	for i, _ := range ans {
		ans[i] -= 1
	}

	for i, v := range nums {
		for len(stack) > 0 && v > nums[stack[len(stack)-1]] {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans[popped] = v
		}
		stack = append(stack, i)
	}

	// stack cleanup
	for _, v := range nums {
		for len(stack) > 0 && v > nums[stack[len(stack)-1]] {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if ans[popped] != -1 {
				continue
			}
			ans[popped] = v
		}
	}

	return ans
}

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(nextGreaterElements2(nums))
}
