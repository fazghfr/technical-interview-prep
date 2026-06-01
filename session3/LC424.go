package main

import (
	"fmt"
)

// AABABBA

// A -> M 1, L 1, -> 1
// AA -> M2, L2
// AAB -> M2, L3, D1, -> 3
// AABA -> M3, L4, D1 -> 4
// AABAB -> M3, L5, D2 -> CANT FLIP ALL, SO SKIP

// VALID WINDOW : WHEN DIFF IS <= K
// ON VALID WINDOW : MAX OF GLOBAL OR DIFF

func characterReplacement(s string, k int) int {
	m := make(map[byte]int)
	windowLen := 0
	currMaxFreq := 0
	left := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		currMaxFreq = max(currMaxFreq, m[s[i]])
		windowLen = i - left + 1
		for windowLen-currMaxFreq > k {
			m[s[left]]--
			left++
			windowLen = i - left + 1
		}
		ans = max(ans, windowLen)
	}
	return ans
}

func main() {
	// fmt.Println(characterReplacement("AABABBA", 1))
	// fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("ABAA", 0))
}
