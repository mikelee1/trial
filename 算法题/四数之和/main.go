package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	var twoAddIndexs = map[int][][]int{}
	var record = map[string]bool{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			tmp := nums[i] + nums[j]
			if _, ok := twoAddIndexs[tmp]; ok {
				twoAddIndexs[tmp] = append(twoAddIndexs[tmp], []int{i, j})
				continue
			}
			twoAddIndexs[tmp] = [][]int{{i, j}}
		}
	}
	for k, v := range twoAddIndexs {
		if v1, ok := twoAddIndexs[target-k]; ok {
			for _, x := range v {
				for _, y := range v1 {
					if x[0] == y[0] || x[0] == y[1] || x[1] == y[0] || x[1] == y[1] {
						continue
					}
					tmp := []int{nums[x[0]], nums[x[1]], nums[y[0]], nums[y[1]]}
					sort.Ints(tmp)
					key := fmt.Sprintf("%d-%d-%d", tmp[0], tmp[1], tmp[2])
					if _, ok := record[key]; ok {
						continue
					}
					record[key] = true
					res = append(res, tmp)
				}
			}
		}
	}
	return res
}


