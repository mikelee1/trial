package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 2))
}

func uniquePaths(m int, n int) int {
	dp := map[int]map[int]int{}
	for k := 0; k < m; k++ {
		dp[k] = map[int]int{}
	}

	for k := 0; k < m; k++ {
		dp[k][0] = 1
	}

	for k := 0; k < n; k++ {
		dp[0][k] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
