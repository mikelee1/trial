package main

import (
	"fmt"
	"time"
)

func main() {
	go test()
	time.Sleep(1*time.Second)
}
// 只要main函数没有退出，test里的协程就会全部产生
func test() {
	for k, _ := range [10]int{} {
		go func() {
			time.Sleep(500*time.Millisecond)
			fmt.Println("k: ", k)
		}()
	}
	fmt.Println("pass test")
}
