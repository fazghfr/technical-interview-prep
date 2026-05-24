package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	myS := strings.Fields(s)
	newS := make([]string, len(myS))
	iterator := len(newS) - 1
	for _, v := range myS {
		newS[iterator] = v
		iterator--
	}

	return strings.Join(newS, " ")
}

func main() {
	fmt.Println(reverseWords("This is the end"))
	fmt.Println(reverseWords(" Reversing words Should work here "))
	fmt.Println(reverseWords("a good   example"))
}
