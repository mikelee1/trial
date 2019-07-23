package main

import "fmt"

var nums = []int{1, 2, 3, 1, 1}

func zhongshu() int {
	if len(nums) == 0 {
		return 0
	}

	i, res := 1, nums[0]

	for _, v := range nums[1:] {
		if v != res {
			i--
		} else {
			i++
		}

		if i <= 0 {
			res = v
			i = 1
		}
	}

	return res

}
func main() {
	a := zhongshu()
	fmt.Println(a)
}
