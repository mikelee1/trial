package main

import (
	"fmt"
)

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	sh, th := 0, 0
	for sh < len(s) && th < len(t) {
		if s[sh] == t[th] {
			sh++
		}
		th++
	}
	return sh == len(s)
}
