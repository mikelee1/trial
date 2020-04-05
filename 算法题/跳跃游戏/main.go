package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	steps := nums[0]
	totalSteps := nums[0]
	var i = 1
	for i <= totalSteps && i < len(nums) {
		steps--
		if steps < nums[i] {
			totalSteps += nums[i] - steps
			steps = nums[i]
		}
		i++
	}
	return totalSteps >= len(nums)-1
}
