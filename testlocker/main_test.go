package main

import (
	"testing"
	"sync"
	"time"
	"fmt"
)

var lock sync.Mutex
var count = make(map[string]int)

func Test_Main(t *testing.T) {
	handler("yiniaji")
	handler("ernianji")
	handler("sannianji")
	handler("sinianji")
	wg.Wait()
}

var wg sync.WaitGroup

func handler(nianji string) {
	count[nianji] = 0
	//var closec = make(chan interface{},1)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		lock.Lock()
		for range []int{1, 2, 3} {

			time.Sleep(3 * time.Second)
			if count[nianji] > 0 {
				break
			}
			fmt.Println(count)
			fmt.Println(nianji)
			count[nianji]++
			fmt.Println(count)
		}
		lock.Unlock()
	}(&wg)
}
