package main

import (
	"fmt"
	"math"
)

func main() {
	a := 70
	i := 0
	b := 0
	var bpre int
	for a != b {
		if b > a {
			fmt.Println(binarysearch(bpre, b, a))
			return
		} else {
			bpre = b
			b = int(math.Pow(float64(2), float64(i)))
			i++
		}
	}
	fmt.Println(b)
}

func binarysearch(bpre, b, a int) int {
	var mid int
	for bpre < b {
		fmt.Println(bpre, b, mid)
		mid = (bpre + b) / 2
		if mid == a {
			break
		}
		if mid > a {
			b = mid
		} else {
			bpre = mid
		}
	}
	return mid
}
