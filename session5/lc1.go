package main

import "fmt"

// 1. store map[target - v] with value of that index
// 2. so that map[x] -> will give us index of vth index
//
//

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		m[target-v] = i
	}
	ans := make([]int, 2)

	for i, v := range nums {
		if _, ok := m[v]; ok {
			ans[0] = i
			ans[1] = m[v]
		}
	}
	return ans
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
