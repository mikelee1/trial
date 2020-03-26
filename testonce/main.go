package main

import (
	"sync"
	"fmt"
	"time"
)

var (
	once = sync.Once{}
	a    int
	aSec time.Duration
)

func main() {
	aSec = time.Second
	fmt.Println(aSec)
	go func() {
		for {
			aa := getA()
			fmt.Println(aa)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			aa := getA()
			fmt.Println(aa)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			aa := getA()
			fmt.Println(aa)
			time.Sleep(time.Second)
		}
	}()

	select {}
}

func getA() int {
	once.Do(func() {
		a = 1
		fmt.Println("once")
	})
	return a
}
