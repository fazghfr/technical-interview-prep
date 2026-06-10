package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// monotonic stack on nums2 with hashmaps
	var stack []int
	m := make(map[int]int)

	for _, v := range nums2 {
		for len(stack) > 0 && v > stack[len(stack)-1] {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			m[popped] = v
		}
		stack = append(stack, v)
	}

	ans := make([]int, len(nums1))
	for i, v := range nums1 {
		if _, ok := m[v]; ok {
			ans[i] = m[v]
		} else {
			ans[i] = -1
		}
	}

	return ans
}

func main() {
	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
