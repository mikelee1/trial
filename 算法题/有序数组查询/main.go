package main

import (
	"fmt"
)

func main() {
	a := searchRange([]int{1, 2, 2, 2, 3, 3, 4, 4, 4, 5}, 5)
	fmt.Println(a)
}

func searchRange(nums []int, target int) []int {
	leng := len(nums)
	var first, last = 0, leng-1
	mid := 0
	for first < last {
		mid = (first + last) / 2
		if nums[mid] > target {
			last = mid - 1
		} else if nums[mid] < target {
			first = mid + 1
		} else {
			break
		}
	}
	fmt.Println(first, last, mid)
	if first == last {
		return []int{first, last}
	}

	if first > last {
		return []int{-1, -1}
	}

	i := mid - 1
	j := mid + 1
	first = mid
	last = mid
	for ; i >= 0 || j < leng; i, j = i-1, j+1 {
		if i >= 0 && nums[i] == target {
			first -= 1
		}
		if j < leng && nums[j] == target {
			last += 1
		}
	}
	return []int{first, last}
}
