package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
}

// 有序数组nums
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
