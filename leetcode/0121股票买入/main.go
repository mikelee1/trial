package main

import (
	"fmt"
	"github.com/hashicorp/consul/lib"
)

func main() {
	lists := []int{
		7,1,5,3,6,4,
	}
	lists = []int{
		7,6,4,3,1,
	}

	fmt.Println(lists)
	FindMaxProfit(lists)
	
}

func FindMaxProfit(nums []int) int {
	var res = 0

	for k,v := range nums{
		for _,v1 := range nums[k:]{
			res = lib.MaxInt(v1-v,res)
		}
	}
	fmt.Println(res)
	return res
}

func FindMaxProfit1(nums []int) int {
	var sum = 0
	ans := 0
	for _,v := range nums{
		if sum > 0{
			sum += v
		}else{
			sum = v
		}
		ans = lib.MaxInt(ans,sum)
	}
	fmt.Println(ans)
	return ans
}