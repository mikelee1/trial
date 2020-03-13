package main

import (
	"testing"
	"fmt"
)

func main() {

}

func Test_Scan1(t *testing.T) {
	a := 0
	b := 0
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		} else {
			fmt.Printf("%d\n", a+b)
		}
	}
}

func Test_Scan2(t *testing.T) {
	n := 0
	ans := 0

	nn, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(nn, err)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := 0
			fmt.Scan(&x)
			ans = ans + x
		}
	}
	fmt.Printf("%d\n", ans)
}
