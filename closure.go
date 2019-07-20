package main

import "fmt"

func main() {
	f := show()
	fmt.Println(f())
	fmt.Println(f())
	f = show()
	fmt.Println(f())
}

func show() func() int {
	x := 0
	return func() int {
		x++
		return x * x
	}
}