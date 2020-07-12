package testcover

import (
	"math"
	"fmt"
)

type BB struct {
	Name string
}

func (a BB) Calculate() {
	fmt.Println(a.Name)
}

func Area(p bool) float64 {
	var r float64 = 1
	if p {
		return math.Pi * r * r
	}
	return math.Pi * r * r
}

func Length() float64 {
	var r float64 = 1
	return math.Pi * 2 * r
}

func Area1() float64 {
	var r float64 = 1
	return math.Pi * r * r
}

func Area2() float64 {
	var r float64 = 1
	return math.Pi * r * r
}
