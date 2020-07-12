package main

import "fmt"

func main() {
	var formatStr = "爬%d个台阶有%d种方式\n"
	fmt.Printf(formatStr, 3, climbStairs(3))
	fmt.Printf(formatStr, 4, climbStairs(4))
	fmt.Printf(formatStr, 5, climbStairs(5))
	//爬3个台阶有3种方式
	//爬4个台阶有5种方式
	//爬5个台阶有8种方式
}

func climbStairs(n int) (dp int) {
	if n <= 2 {
		return n
	}
	prePre, pre := 1, 2
	for i := 2; i < n; i++ {
		prePre, pre = pre, pre+prePre
	}
	return pre
}
