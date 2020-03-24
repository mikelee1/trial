package main

import "fmt"

func main() {
	fmt.Println(climbStairs(3))
}

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	dp := 0
	pre := 2
	prepre := 1
	for i := 2; i < n; i++ {
		dp = pre + prepre
		prepre = pre
		pre = dp
	}
	return dp
}
