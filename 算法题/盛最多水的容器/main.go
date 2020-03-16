package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	maxA := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			temp := compute(height[i], height[j], j-i)
			if temp > maxA {
				maxA = temp
			}
		}
	}
	return maxA
}

func compute(left, right, long int) int {
	if left < right {
		return long * left
	}
	return long * right
}
