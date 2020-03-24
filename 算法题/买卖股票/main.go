package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := 0
	min := prices[0]
	pre := 0
	for i := 1; i < len(prices); i++ {
		dp = max(pre, prices[i]-min)
		if prices[i] < min {
			min = prices[i]
		}
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
