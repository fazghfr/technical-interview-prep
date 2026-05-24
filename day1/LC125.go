package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) bool {
    str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	// cleaning the string of non alphanumeric character
	var newStr []rune
	for _, v := range str {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			newStr = append(newStr, rune(v))
		} else {
			continue
		}
	}

	left :=  0
	right := len(newStr) - 1
	for left < right {
		if newStr[left] != newStr[right] {
			return false
		}
		right--
		left++
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("malam3514313"))
	fmt.Println(isPalindrome(" Mala62371836821m"))
}
