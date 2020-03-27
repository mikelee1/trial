package main

import (
	"fmt"
	"myproj/try/teststruct/Inner"
)

func main() {
	a := A{

	}
	a.Print()
	a.B.Print()
}

type A struct {
	Name string
	C
}

type C struct {
	inner.B
}

func (c C) Print() {
	fmt.Println("c")
}
