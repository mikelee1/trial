package main

import (
	"fmt"
	"strings"
)

func main() {
	c := "杭州市"
	a := strings.Split(c, "")
	if a[len(a)-1] == "市" {
		fmt.Println(Substr(c, 0, len(a)-1))
	} else {
		fmt.Println(Substr(c, 0, 2))
	}
}

// Substr returns the substr from start to length.
func Substr(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}
