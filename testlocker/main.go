package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.RWMutex
//读写锁
func main() {

	fmt.Println("The lock is locked.()")
	mutex.Lock()

	go func() {
		for range [10]int{} {
			go read()
		}
	}()
	time.Sleep(time.Second)

	fmt.Println("The lock is unlocked.()")
	mutex.Unlock()

	//休息一会,等待打印结果
	time.Sleep(time.Second)
}

func read() {
	fmt.Println("The lock is rlocked.()")
	mutex.RLock()
	defer mutex.RUnlock()
	fmt.Println("The lock is runlock.()")
}
