package main

import "fmt"

func main() {
	var a, b []int
	realLoad(&a)
	fakeLoad(b)
	fmt.Println(a)
	fmt.Println(b)
}

func realLoad(aa *[]int) {
	*aa = append(*aa, []int{1, 2}...)
}

func fakeLoad(aa []int) {
	aa = []int{1, 2}
}
