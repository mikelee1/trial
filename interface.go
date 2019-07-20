// _Interfaces_ are named collections of method
// signatures.

package main

import "fmt"
import "math"

// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// For our example we'll implement this interface on
// `rect1` and `circle` types.
type rect1 struct {
	width, height float64
}
type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect1`s.
func (r rect1) area() float64 {
	return r.width * r.height
}
func (r rect1) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for `circle`s.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
//
//func main() {
//	r := rect1{width: 3, height: 4}
//	c := circle{radius: 5}
//
//	// The `circle` and `rect1` struct types both
//	// implement the `geometry` interface so we can use
//	// instances of
//	// these structs as arguments to `measure`.
//	measure(r)
//	measure(c)
//}

type T1 interface {
	Add1() string
	Minus() string
}
type T2 interface {
	Add1() string
	Minus() string
}

type Data struct {
	A T2
}
func (d Data)Add1() string {
	fmt.Println("a")
	return ""
}

func (d Data)Minus() string {
	fmt.Println("a")
	return ""
}

type Test1 struct {
	A Data
}
//
//func main()  {
//	t := Test1{}
//	t.A.Add1()
//}

type Manager interface {
	PaySalary()
}

type Staff struct {
	Grade int
}

func (s Staff)PaySalary()  {
	switch s.Grade {
	case 1:
		fmt.Println("2w")
	case 2:
		fmt.Println("1W")
	}
}

func main()  {
	var s Manager = Staff{1}
	s.PaySalary()

	s = Staff{2}
	s.PaySalary()
}