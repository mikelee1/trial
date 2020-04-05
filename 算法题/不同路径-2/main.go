package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0},
		{1, 1},
		{0, 0},
	}))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	length := len(obstacleGrid)
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	dp := map[int]map[int]int{}

	for i := 0; i < length; i++ {
		dp[i] = map[int]int{}
	}
	dp[0][0] = 1

	for i := 1; i < length; i++ {
		if obstacleGrid[i][0] == 1 || dp[i-1][0] == 0 {
			dp[i][0] = 0
			continue
		}
		dp[i][0] = 1
	}

	for i := 1; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 1 || dp[0][i-1] == 0 {
			dp[0][i] = 0
			continue
		}
		dp[0][i] = 1
	}
	//fmt.Println(dp)
	var i, j = len(obstacleGrid), len(obstacleGrid[0])
	for i = 1; i < length; i++ {
		for j = 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	//fmt.Println(dp)
	return dp[i-1][j-1]
}
