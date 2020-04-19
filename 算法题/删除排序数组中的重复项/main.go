package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2, 3, 4, 5}))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var old = nums[0] - 1
	for k := 0; k < len(nums); {
		if old == nums[k] {
			nums = append(nums[:k], nums[k+1:]...)
		} else {
			old = nums[k]
			k++
		}
	}
	return len(nums)
}
