package main

import (
	"fmt"
	"github.com/hashicorp/consul/lib"
)

func main() {
	lists := []int{
		-2, 1, -3, 4, -1, 2, 1, -5, 4,
	}
	FindMaxSum(lists)
	FindMaxSum1(lists)
}

func FindMaxSum(nums []int) {
	if len(nums) == 1 {
		fmt.Println(nums[0])
		return
	}
	//var head, tail int
	ressum := nums[0]
	for i := 0; i < len(nums); i++ {
		newsum := nums[i]
		maxsum := nums[i]
		//newhead := i
		//newtail := i
		for j := i + 1; j < len(nums); j++ {
			if newsum+nums[j] >= maxsum {
				maxsum = newsum + nums[j]
				//newtail = j
			}
			newsum = newsum + nums[j]
		}
		if maxsum > ressum {
			ressum = maxsum
			//head = newhead
			//tail = newtail
		}
	}
	//fmt.Println(head,tail)
	fmt.Println(ressum)
	return
	//return l[head:tail+1]
}

func FindMaxSum1(nums []int) int {
	var sum = 0
	ans := nums[0]
	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		ans = lib.MaxInt(ans, sum)
	}
	fmt.Println(ans)
	return ans
}
