package main

import "fmt"

func finalPrices(prices []int) []int {
	var stack []int

	for i, v := range prices {
		for len(stack) > 0 && v <= prices[stack[len(stack)-1]] {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			prices[popped] -= v
		}
		stack = append(stack, i)
	}
	return prices
}

func main() {
	prices := []int{8, 4, 6, 2, 3}
	fmt.Println(finalPrices(prices))
}
