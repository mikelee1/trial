package main

import (
	"fmt"
)

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 2, 5},
		{3, 2, 1},
	}))
}

func minPathSum(grid [][]int) int {
	//先初始化二维数组变量
	dp := make([][]int, len(grid))
	for k := 0; k < len(grid); k++ {
		dp[k] = make([]int, len(grid[k]))
	}

	//准备好边界值
	dp[0][0] = grid[0][0]
	for j := 1; j < len(dp[0]); j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < len(dp); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	//开始动态规划
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = min(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
