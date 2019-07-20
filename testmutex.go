package main

import (
	"sync"
	"fmt"
	"time"
)

var lock11 *sync.Mutex

func init()  {
	lock11 = &sync.Mutex{}
}

func add(i *int)  {
	time.Sleep(1*time.Second)
	lock11.Lock()
	defer lock11.Unlock()
	*i++
}

func main() {
	i := 0
	for range [200]int{}{
		go add(&i)
	}
	time.Sleep(2*time.Second)
	fmt.Println(i)
}
