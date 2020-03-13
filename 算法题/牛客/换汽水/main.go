package main

import "fmt"

func main() {
	n := 0
	res := ""

	for range [10]int{} {
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		res += fmt.Sprintf("%d\n", MaxBottle(n))
	}
	fmt.Printf("%s", res)
}

func MaxBottle(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	remain := n % 3
	newBottle := (n - remain) / 3

	return newBottle + MaxBottle(remain+newBottle)
}
