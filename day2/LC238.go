package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	arrLeft := make([]int, len(nums))
	for i, v := range nums {
		if i == 0 {
			arrLeft[i] = v
			continue
		}
		arrLeft[i] = arrLeft[i-1] * v
	}
	arrRight := make([]int, len(nums))
	for i:=len(nums)-1; i >= 0; i-- {
		v := nums[i]
		if i == len(nums) -1 {
			arrRight[i] = v
			continue
		}
		arrRight[i] = arrRight[i+1] * v
	}

	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			ans[i] = arrRight[i+1]
			continue
		}
		if i == len(nums) - 1 {
			ans[i] = arrLeft[i-1]
			continue
		}

		ans[i] = arrLeft[i-1] * arrRight[i+1]
	}
	return ans
}

func main() {
	arr := []int{1,2,3,4}
	ans := productExceptSelf(arr)
	fmt.Println(ans)
}
