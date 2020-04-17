package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 3, 2, 1, 5, 7, 4, 6, 5, 8, 9, 2}
	sort.Ints(a)
	fmt.Println(twoSum(a, 6))
}

// 有序数组nums，找到两个数加起来等于target
func twoSum(nums []int, target int) []int {

	var head, tail = 0, len(nums)-1
	for head < tail {
		a := nums[head] + nums[tail]
		switch {
		case a == target:
			return []int{head + 1, tail + 1}
		case a > target:
			tail--
		default:
			head++
		}
	}
	return []int{-1, -1}
}
