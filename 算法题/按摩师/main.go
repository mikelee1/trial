package main

import "fmt"

func main() {
	fmt.Println(massage([]int{8, 3, 1, 4, 9, 8, 1, 5, 7, 8, 2, 3}))
}

func massage(nums []int) int {
	n := len(nums)
	dp := 0
	pre := 0
	prepre := 0
	for i := 0; i < n; i++ {
		dp = max(prepre+nums[i], pre)
		prepre = pre
		pre = dp
	}

	return dp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
