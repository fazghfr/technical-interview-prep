package main

import (
	"fmt"
	"sort"
	"strings"
)

// In other words, return true if one of s1's permutations is the substring of s2.

func checkInclusion(s1 string, s2 string) bool {
	r1 := []rune(s1)
	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})
	r2 := []rune(s2)

	windowSize := len(r1)

	for left := 0; left+windowSize <= len(s2); left++ {
		tempR := make([]rune, windowSize)
		copy(tempR, r2[left:left+windowSize])
		sort.Slice(tempR, func(i, j int) bool {
			return tempR[i] < tempR[j]
		})
		if strings.Compare(string(tempR), string(r1)) == 0 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}
