package main

import "fmt"

func main() {
	fmt.Println(maxYue(1992,900))
}

func maxYue(a, b int) int {
	if a == 1 || b == 1 {
		return 1
	}
	if a%b == 0 {
		return b
	}
	if b%a == 0 {
		return a
	}
	if a > b {
		a, b = b, a
	}
	return maxYue(a, b-a)
}
