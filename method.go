package main

import (
	"fmt"
	"time"
)

type rect struct {
	width, height int
}

// This `area` method has a _receiver type_ of `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value
// receiver types. Here's an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	rp := &r
	t := time.Now()
	// Here we call the 2 methods defined for our struct.
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
	fmt.Println(time.Since(t))
	t1 := time.Now()

	// Go automatically handles conversion between values
	// and pointers for method calls. You may want to use
	// a pointer receiver type to avoid copying on method
	// calls or to allow the method to mutate the
	// receiving struct.

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
	fmt.Println(time.Since(t1))
}
