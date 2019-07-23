package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	tm2 := tm1.AddDate(0, 0, 1)
	fmt.Println(tm1, tm2)
}
