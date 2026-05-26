package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	temp := make(map[string][]string)

    for _, str := range strs {
    	var myR []rune
    	for _,v  := range str {
     		myR = append(myR, v)
     	}
      	sort.Slice(myR, func(i, j int) bool {
       		return myR[i] <= myR[j]
       })
      	temp[string(myR)] = append(temp[string(myR)], str)
    }

    var ans [][]string
    for _, v := range temp {
   		ans = append(ans, v)
    }

    return ans
}

func main() {
	testStr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	//
	fmt.Println(groupAnagrams(testStr))
}
