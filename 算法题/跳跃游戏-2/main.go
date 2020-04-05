package main

import "fmt"

func main() {
	fmt.Println(jump([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}))
	//fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
func jump(nums []int) int {
	var i = 0
	var steps = 1
	var remainSteps = nums[0]
	if len(nums) == 1 {
		return 0
	}
	for i < len(nums) {
		if i+remainSteps >= len(nums)-1 {
			return steps
		}
		subMax := 0
		maxJ := 1
		pre := 0
		for j := 1; j <= remainSteps; j++ {
			if nums[i+j]+j >= pre {
				pre = nums[i+j] + j
				subMax = nums[i+j]
				maxJ = j
			}
		}

		remainSteps = subMax
		i = i + maxJ
		steps++
	}
	return steps
}
