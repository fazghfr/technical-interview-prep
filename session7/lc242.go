package main

import "fmt"

func isAnagram(s string, t string) bool {
	// frequency map
	if len(s) != len(t) {
		return false
	}
	freq := make(map[rune]int)

	for _, v := range s {
		freq[v]++
	}

	for _, v := range t {
		if _, ok := freq[v]; ok {
			freq[v]--
		} else {
			return false
		}
	}

	for _, v := range freq {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "ana"
	t := "naa"
	fmt.Println(isAnagram(s, t))
}
