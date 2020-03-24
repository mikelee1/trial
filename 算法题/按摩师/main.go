package main

import "fmt"

func main() {
	fmt.Println(massage([]int{2, 7, 9, 3, 1, 4, 5, 2, 7, 2, 1}))
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
