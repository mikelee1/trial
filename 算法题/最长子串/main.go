package main

import (
	"fmt"
)

var str = "aabaab!bb"

func main() {
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(ss string) int {
	var head, tail = 0, 0
	var max = 0
	var m = make(map[byte]int)
	for k, s := range []byte(ss) {
		if index, ok := m[s]; ok {
			m[s] = k
			head = Max(index+1, head)
		} else {
			m[s] = k
		}
		tail = k + 1
		max = Max(max, tail-head)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
