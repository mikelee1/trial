package main

import "fmt"

var ab = [][]int{
	[]int{1, 4, 7, 11, 15},
	[]int{2, 5, 8, 12, 19},
	[]int{3, 6, 9, 16, 22},
	[]int{10, 13, 14, 17, 24},
	[]int{18, 21, 23, 26, 30},
}

func main() {
	fmt.Println(ab)
	for _, v := range ab {
		fmt.Println(v)
	}
}
