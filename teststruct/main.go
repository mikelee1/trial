package main

import (
	"myproj/try/teststruct/Inner"
	"fmt"
)

func main() {
	t111 := Test111{Test222:Inner.Test222{Address:"1sdf"}}
	fmt.Println(t111)
}

type Test111 struct {
	Name string
	Inner.Test222
}



