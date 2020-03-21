package main

import "fmt"

func main() {
	fmt.Println(canMeasureWater(0, 0, 0))
}

func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}
	if z > x+y || z < 0 {
		return false
	}
	if x*y == 0 {
		if x+y != z {
			return false
		}
		return true
	}
	if z%maxYue(x, y) == 0 {
		return true
	}
	return false
}

func subMinus(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func change(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func maxYue(a, b int) int {
	if a == 1 || b == 1 {
		return 1
	}
	if a%b == 0 {
		return b
	}
	if b%a == 0 {
		return a
	}
	a, b = change(a, b)
	return maxYue(a, subMinus(a, b))
}
