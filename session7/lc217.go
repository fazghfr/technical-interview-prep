package main

import "fmt"

func containsDuplicate(nums []int) bool {
	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++
		if freq[v] > 1 {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
}
