package main

import "fmt"

func carPooling(trips [][]int, capacity int) bool {
	// last stop
	lastStop := 0
	for _, v := range trips {
		lastStop = max(lastStop, v[2])
	}

	diffArray := make([]int, lastStop+1)
	for _, v := range trips {
		diffArray[v[1]] += v[0]
		diffArray[v[2]] -= v[0]
	}

	prefixSums := make([]int, len(diffArray))
	for i, v := range diffArray {
		if i == 0 {
			prefixSums[i] += v
		} else {
			prefixSums[i] += prefixSums[i-1] + v
		}

		if prefixSums[i] > capacity {
			return false
		}
	}
	return true
}

// [[2,1,5],[3,3,7]], capacity = 5
// [[2,1,5],[3,3,7]], capacity = 4
// 1 2 3 4 5 6 7
// 2 2 5 5 3 3 0

func main() {
	trips := [][]int{
		{2, 1, 5},
		{3, 3, 7},
	}

	capacity := 5
	fmt.Println(carPooling(trips, capacity))
}
