package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type Order struct {
	Money    float64
	Quantity int

	PrePrice   float64
	FinalPrice float64
}

func main() {
	order := &Order{
		Money:    2,
		Quantity: 3,
	}
	order.TransOrder()
	fmt.Println(order.PrePrice, order.FinalPrice)

	order = &Order{
		Money:    1,
		Quantity: 80,
	}
	order.TransOrder()
	fmt.Println(order.PrePrice, order.FinalPrice)

}

func (order *Order) TransOrder() {
	order.PrePrice, _ = decimal.NewFromFloat(order.Money).Div(decimal.NewFromFloat(float64(order.Quantity))).Truncate(2).Float64()
	order.FinalPrice, _ = decimal.NewFromFloat(order.Money).Sub(decimal.NewFromFloat(order.PrePrice).Mul(decimal.NewFromFloat(float64(order.Quantity - 1)))).Truncate(2).Float64()
}
