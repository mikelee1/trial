package main

import (
	"time"
	"fmt"
)

func main() {
	t,err := time.Parse("20060102","199201041")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(t)
}
