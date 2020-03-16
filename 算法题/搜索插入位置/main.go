package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1}, 0))
}

func searchInsert(nums []int, target int) int {
	head, tail := 0, len(nums)-1
	mid := 0
	for head <= tail {
		mid = (head + tail) / 2
		if nums[mid] > target {
			tail = mid - 1
			continue
		}
		if nums[mid] < target {
			head = mid + 1
			continue
		}
		return mid
	}
	return head
}
