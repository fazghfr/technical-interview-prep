package main

import "fmt"

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]

	for i, v := range prices {
		if i == 0 {
			continue
		}
		maxProfit = max(maxProfit, v-minPrice)
		minPrice = min(minPrice, v)
	}
	return maxProfit
}

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(arr))
}
