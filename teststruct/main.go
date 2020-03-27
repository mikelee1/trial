package main

import (
	"fmt"
)

func main() {
	a := A{

	}
	a.Print()
}

type A struct {
	Name string
	C
}

type C struct{}

func (c C) Print() {
	fmt.Println("c")
}
