package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 100, 1, 100, 1}))
}

func minCostClimbingStairs(cost []int) int {
	switch len(cost) {
	case 0:
		return 0
	case 1:
		return 0
	case 2:
		return min(cost[0], cost[1])
	default:
		dp := 0
		pre := cost[1]
		prepre := cost[0]
		for i := 2; i < len(cost); i++ {
			dp = min(pre+cost[i], prepre+cost[i])
			prepre = pre
			pre = dp
		}
		return min(pre, prepre)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
