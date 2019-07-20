package main

import (
	"fmt"
	"time"
)

type AA struct {
	Value string
}

func main()  {
	//A1 := AA{"a1"}
	//A2 := AA{"a2"}
	//a := []AA{A1,A2}
	//for k,_ := range a{
	//	a[k].Value = "a3"
	//}
	//fmt.Println(a)
	forgoroutine()
	time.Sleep(5*time.Second)
}

func forgoroutine()  {
	for _, node := range []int{1,2,3,4} {
		go func() {
			fmt.Println(node)
		}()
	}
}