package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(-121))
}


func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
