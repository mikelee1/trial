package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumNums(4))
}
func sumNums(n int) int {
	if n == 1 {
		return 1
	}
	return n + sumNums(n-1)
}
