package main

import (
	"fmt"
)

func Add(a,b int) int {
	return a+b
}

func Wrapper(f func(a,b int) int) func(a,b int) int {
	return func(a, b int) int {
		fmt.Println("add")
		return f(a,b)
	}
}

func main() {
	 f := Wrapper(Add)
	 b := f(1,2)
	 fmt.Println(b)
}
