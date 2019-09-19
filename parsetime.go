package main

import (
	"fmt"
	"time"
)


func main() {
	//t, err := time.Parse("20060102", "199201041")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(t)
	a := "2019-09-18T16:00:00.000Z"
	b,_ := time.Parse(time.RFC3339,a)
	fmt.Println(b.Format("2006"))
}
