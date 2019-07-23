package main

import "fmt"

func main() {
	one := 1
	changevalue(one)
	fmt.Println(one)

	two := 1
	changepointer(&two)
	fmt.Println(two)
}

func changevalue(a int) {
	a = 0
}

func changepointer(a *int) {
	*a = 0
}
