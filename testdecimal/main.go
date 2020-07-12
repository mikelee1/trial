package main

import (
	"github.com/shopspring/decimal"
	"fmt"
)

func main() {
	a := float64(0.05)
	b := float64(0.01)
	a1 := decimal.NewFromFloat(a)
	b1 := decimal.NewFromFloat(b)

	fmt.Println(a1.Round(1))

	fmt.Println(a1.Mul(b1))
	fmt.Println(a1.Sub(b1).DivRound(decimal.NewFromFloat(3), 3).Value())
	fmt.Println(float64(a-b) / 3.0)
}
