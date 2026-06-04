package main

import "fmt"

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n+1)
	for _, booking := range bookings {
		left := booking[0]
		right := booking[1]
		seats := booking[2]
		diff[left] += seats
		if right+1 < len(diff) {
			diff[right+1] -= seats
		}
	}

	ans := make([]int, n+1)
	for i, v := range diff {
		if i == 0 {
			ans[i] = v
			continue
		}
		ans[i] += v + ans[i-1]
	}

	return ans[1:]
}

// idea deduction
//
// brute force approach
//
// for every bookings, write from left to right
//

func main() {
	bookings := [][]int{
		{1, 2, 10},
		{2, 3, 20},
		{2, 5, 25},
	}

	fmt.Println(corpFlightBookings(bookings, 5))
}
