package main

import (
	"fmt"
	"strconv"
	"strings"
	"math"
)

func reverse(x int) int {
	isNegative := false
	if x < 0 {isNegative = true}
	xstr := strconv.Itoa(x)
	if isNegative{
		xstr = strings.Trim(xstr, "-")
	}

	ans := make([]byte, len(xstr))
	it := 0
	for i := len(xstr) - 1; i>=0;i-- {
		ans[it] = xstr[i]
		it++
	}

	ansInt, _ := strconv.Atoi(string(ans))

	if isNegative {
		ansInt = ansInt * -1
		if ansInt < int(math.Pow(2, 31)) *-1 {
			return 0
		}
	}
	if ansInt > int(math.Pow(2, 31) - 1) {
		return 0
	}
	return ansInt
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(-67))
	fmt.Println(reverse(67))
}
