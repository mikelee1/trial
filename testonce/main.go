package main

import (
	"sync"
	"fmt"
	"time"
)

var (
	once = sync.Once{}
	a int
)



func main() {
	go func() {
		for {
			aa := GetA()
			fmt.Println(aa)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			aa := GetA()
			fmt.Println(aa)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			aa := GetA()
			fmt.Println(aa)
			time.Sleep(time.Second)
		}
	}()

	select {

	}
}

func GetA() int {
	once.Do(func() {
		a = 1
		fmt.Println("once")
	})
	return a
}