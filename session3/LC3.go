package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	left := 0
	maxLen := -1

	for right := 0; right < len(s); right++ {
		seen[s[right]]++
		for seen[s[right]] != 1 {
			seen[s[left]]--
			left++
		}
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
