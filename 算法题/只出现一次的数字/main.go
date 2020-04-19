package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{3, 2, 2, 3, 4}))
}

func singleNumber(nums []int) int {
	var res int
	for _, v := range nums {
		res = res ^ v
	}
	return res
}
