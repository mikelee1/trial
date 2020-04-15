package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	mIndex := map[int][]int{}
	for k, n := range nums {
		_, ok := mIndex[n]
		if ok {
			mIndex[n] = append(mIndex[n], k)
		} else {
			mIndex[n] = []int{k}
		}
	}
	for _,n1 := range nums {
		v, ok := mIndex[target-n1]
		if ok {
			if target-n1 != n1 {
				return []int{mIndex[n1][0], v[0]}
			}
			if len(v) < 2 {
				continue
			}
			return []int{v[0], v[1]}
		}
	}
	return []int{-1, -1}
}
