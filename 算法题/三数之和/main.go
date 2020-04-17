package main

import (
	"fmt"
)

func main() {
	fmt.Println(threeSum([]int{0,0,0,0,0,0,0}))
}

// 三数之和
func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	var recorded = map[string]bool{}
	var singleIndexs = map[int][]int{}
	for k, v := range nums {
		if _, ok := singleIndexs[v]; ok {
			if len(singleIndexs[v])<3{
				singleIndexs[v] = append(singleIndexs[v], k)
			}
			continue
		}
		singleIndexs[v] = []int{k}
	}

	for head := 0; head < len(nums)-1; head++ {
		for mid := head + 1; mid < len(nums); mid++ {
			twosum := nums[head] + nums[mid]
			vs, ok := singleIndexs[-twosum]
			if !ok {
				continue
			}
			for _, v := range vs {
				if v == head || v == mid {
					continue
				}
				key := Sort(nums[v], nums[head], nums[mid])
				if _, ok := recorded[key]; !ok {
					recorded[key] = true
					res = append(res, []int{nums[v], nums[head], nums[mid]})
				}
			}
		}
	}

	return res
}

func Sort(a, b, c int) string {
	switch {
	case a >= b && a >= c && b >= c:
		return fmt.Sprintf("%d-%d", a, b)
	case a >= b && a >= c && b <= c:
		return fmt.Sprintf("%d-%d", a, c)
	case b >= a && b >= c && a >= c:
		return fmt.Sprintf("%d-%d", b, a)
	case b >= a && b >= c && a <= c:
		return fmt.Sprintf("%d-%d", b, c)
	case c >= b && c >= a && b >= a:
		return fmt.Sprintf("%d-%d", c, b)
	default:
		return fmt.Sprintf("%d-%d", c, a)
	}
}
