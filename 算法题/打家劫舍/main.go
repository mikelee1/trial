package main

import "fmt"

func main() {
	fmt.Println(rob([]int{8, 3, 1, 4, 9, 8, 1, 5, 7, 8, 2, 3}))
}

func rob(nums []int) int {
	dp := make([]int, len(nums))
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	default:
		dp[0] = nums[0]
		dp[1] = max(nums[0], nums[1])
		maxRob := nums[0]
		for i := 2; i < len(nums); i++ {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
			maxRob = max(dp[i], maxRob)
		}
		return maxRob
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
