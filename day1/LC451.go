package main

import (
	"fmt"
	"sort"
	"strings"
)

func frequencySort(s string) string {
    m := make(map[rune]int)
    var sb strings.Builder
    for _, v := range s {
    	m[v]++
     	sb.WriteRune(v)
    }

    arrRune := []rune(sb.String())
    sort.Slice(arrRune, func(i, j int) bool {
    		if m[arrRune[i]] == m[arrRune[j]] {
      			return arrRune[i] > arrRune[j]
      		}
            return m[arrRune[i]] > m[arrRune[j]]
    })

    return string(arrRune)
}

func main() {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("loveleetcode"))
}
