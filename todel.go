package main

import "fmt"

func main() {
	a := 11
	switch {
	case a<11:
		fmt.Println("<11")
	case a>11:
		fmt.Println(">11")

	}
	fmt.Println("end")
}
