package main

import "fmt"

var a = []int{1,2,2,3,3}

func singleNumber(nums []int) int {
	var out= 0;
	for i := 0; i < len(nums); i++ {

		out = out ^ nums[i];
	}

	return out;
}

func bitcomp()  {
	fmt.Println(2^3)
}
func main()  {
	fmt.Println(singleNumber(a))
	bitcomp()
}
