package testcover

import "math"

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

func area1() float64 {
	var r float64 = 1
	return math.Pi * r * r
}

func area2() float64 {
	var r float64 = 1
	return math.Pi * r * r
}
