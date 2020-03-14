package main

import (
	"fmt"
)

var (
	a = []int{1,3,5,7,7}
	b = []int{2,4,6,8}
	c = make([]int, 0)
)

func main(){
	res := Combine(a, b)
	fmt.Println(res)
}


func Combine(a, b []int) []int {
	if len(a) == 0 {
		c = append(c, b...)
		return b
	}
	if len(b) == 0 {
		c = append(c, a...)
		fmt.Println(a)
		return a
	}
	if a[0] < b[0] {
		c = append(c, a[0])
		Combine(a[1:], b[:])
	} else {
		c = append(c, b[0])
		Combine(a[:], b[1:])
	}
	return c
}