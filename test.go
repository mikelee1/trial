package main

import (
	"fmt"
	"math"
)

func main() {
	a := 5
	fmt.Println(int(math.Floor(float64(a) * 0.8)))
	b := fmt.Sprintf(" and a.realname like '%%%s%%' ", "李")
	fmt.Println(b)
}
