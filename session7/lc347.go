package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	var keysArr []int
	for keys, _ := range freq {
		keysArr = append(keysArr, keys)
	}

	sort.Slice(keysArr, func(i, j int) bool {
		return freq[keysArr[i]] > freq[keysArr[j]]
	})

	return keysArr[:k]
}

func main() {
	nums := []int{1, 1, 1, 2, 2, -1, -1, -1}
	fmt.Println(topKFrequent(nums, 2))
}

// frequency table
// extract the values into separate array
//
