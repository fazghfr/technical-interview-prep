package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// frequency count
	freq := make(map[rune]int)
	for _, v := range s {
		freq[v]++
	}

	for _, v := range t {
		freq[v]--
	}

	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anjaay", "najaay"))
	fmt.Println(isAnagram("anjaay", "hidup"))
	fmt.Println(isAnagram("nasigoreng", "hidup"))
	fmt.Println(isAnagram("gorengnasi", "nasigerong"))
}
