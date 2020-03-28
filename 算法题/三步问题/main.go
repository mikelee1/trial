package main

import (
	"fmt"
)

func main() {
	fmt.Println(waysToStep(73))
}

func waysToStep(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	}
	var dp = 0
	var preprepre = 1
	var prepre = 2
	var pre = 4
	for i := 4; i < n+1; i++ {
		dp = (pre + prepre + preprepre) % 1000000007
		preprepre = prepre
		prepre = pre
		pre = dp
	}
	return dp
}
