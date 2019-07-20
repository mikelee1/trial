package main

import (
	"time"
	"fmt"
)

func main()  {
	trans(nil)
}

func trans(t interface{})  {
	_,ok := t.(time.Time)
	if ok{
		fmt.Println("is time.time")
	}else{
		fmt.Println("not time.time")
	}
}
