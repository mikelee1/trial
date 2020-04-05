package main

import "fmt"

func main() {
	fmt.Println(divisorGame(3))
}

func divisorGame(N int) bool {
	var dq = map[int]bool{}
	dq[1] = false
	dq[2] = true
	for i := 3; i <= N; i++ {
		for k := 1; k < i; k++ {
			//找到约数，并且该约数对应的i-k为false
			if i%k == 0 && dq[i-k] == false {
				dq[i] = true
			}
		}

	}
	return dq[N]
}
