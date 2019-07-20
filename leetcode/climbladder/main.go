package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(climbLadder(30))
	fmt.Println(time.Now().Sub(start).Nanoseconds())

	//未用map存储
	start2 := time.Now()
	fmt.Println(climbLadder2(30))
	fmt.Println(time.Now().Sub(start2).Nanoseconds())
}

var history = make(map[int]int)

func climbLadder(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if res,ok := history[n];ok{
		return res
	}
	a := climbLadder(n-2)
	history[n-2] = a
	b := climbLadder(n-1)
	history[n-1] = b

	return a+b
}

func climbLadder2(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	a := climbLadder2(n-2)
	b := climbLadder2(n-1)
	return a+b
}
