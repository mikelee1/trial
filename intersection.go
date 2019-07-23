package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 8, 10}
	b := []int{3, 4, 5, 7, 8}
	c := Searchinteruser(a, b)
	fmt.Println(c)
}

func Searchinteruser(a, b []int) []int {
	fmt.Println(a, b)
	ai := 0
	bi := 0
	c := []int{}
	for {
		fmt.Println(ai, bi)
		if ai > len(a)-1 || bi > len(b)-1 {
			break
		}
		if a[ai] == b[bi] {
			c = append(c, a[ai])
			ai++
			bi++
			continue
		}
		if a[ai] > b[bi] {
			bi++
			continue
		} else {
			ai++
		}
	}
	return c
}
