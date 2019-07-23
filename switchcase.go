package main

import "fmt"

func test(a int) {
	switch a {
	case 1:
	case 2:
		fmt.Println(a)
	default:
		fmt.Println("not 1,2")

	}
}

var ch chan interface{}

func test1() {
	select {
	case ch <- struct{}{}:

	}
}

func main() {
	test(1)
	test(2)
	test(3)
	go test1()
}
