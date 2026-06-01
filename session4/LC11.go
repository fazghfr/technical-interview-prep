package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		currWidth := min(height[right], height[left])
		maxArea = max(maxArea, (right-left)*currWidth)

		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

// area calculation is right - left * (minheight of right/left)
// left++ if left is the smaller one
// right-- if right is the smaller one

// the idea is we always shrink to the smaller height to compensate the shrinking width
func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// height := []int{1, 1}
	fmt.Println(maxArea(height))
}
