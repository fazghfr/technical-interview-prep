package main

import "fmt"

func dailyTemperatures(temp []int) []int {
	var myStack []int             // stores the indices
	ans := make([]int, len(temp)) // stores the actual answer

	for i, v := range temp {
		// pop when we find days bigger than top of stack
		for len(myStack) > 0 && v > temp[myStack[len(myStack)-1]] {
			topIdx := myStack[len(myStack)-1]
			myStack = myStack[:len(myStack)-1]

			ans[topIdx] = i - topIdx
		}
		myStack = append(myStack, i)
	}

	return ans
}

func main() {
	temps := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temps))
}
