package main

import "fmt"

func maximumPopulation(logs [][]int) int {
	// last and first year
	lastYear := -1
	firstYear := 2051

	for _, v := range logs {
		lastYear = max(lastYear, v[1])
		firstYear = min(firstYear, v[0])
	}

	// lastyear-firstyear+1 -> the only size that matters.
	//  0 to lastyear is a waste of memory space
	diffArr := make([]int, lastYear-firstYear+1)

	for _, v := range logs {
		writeAdd := v[0] - firstYear
		diffArr[writeAdd] += 1
		writeSub := v[1] - firstYear
		diffArr[writeSub] -= 1
	}

	prefixSum := make([]int, len(diffArr))
	theMaximum := len(prefixSum) * -1
	for i, v := range diffArr {
		if i == 0 {
			prefixSum[i] += v
		} else {
			prefixSum[i] += v + prefixSum[i-1]
		}
		theMaximum = max(prefixSum[i], theMaximum)
	}

	for i, v := range prefixSum {
		if v == theMaximum {
			return i + firstYear
		}
	}
	return -1
}

// [[1993,1999],[2000,2010]]

// 93 94 95 99 200, ...2010
// 1  0  0  -1 +1

func main() {
	trips := [][]int{
		{2000, 2010},
		{1993, 1999},
	}

	fmt.Println(maximumPopulation(trips))
}
