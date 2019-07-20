package main

import (
	"sync"
	"fmt"
	"time"
)

var lock1 sync.RWMutex
var lockvar int

func main()  {

	lock1.Lock()

	closec := make(chan int,1)
	lockvar = 1
	fmt.Println(lockvar)

	go func() {
		time.Sleep(5*time.Second)
		closec<-1
	}()


	for _,v := range []int{1,2,3}{
		go func(v int) {
			lock1.Lock()
			defer lock1.Unlock()
			lockvar++
			fmt.Println(v,lockvar)
		}(v)
	}
	lock1.Unlock()
	time.Sleep(3*time.Second)
	<-closec
}
