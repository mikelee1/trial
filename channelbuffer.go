package main

import (
"fmt"
"sync"
)

var wt sync.WaitGroup

func OutChan(noBufChan <-chan int) {

	//chan关闭后for退出循环，否则会死循环
	for v := range noBufChan {
		fmt.Println(v)
	}

	wt.Done()
}

func InChan(noBufChan chan<- int) {

	for i := 0; i < 5; i++ {
		fmt.Println("write")
		noBufChan <- i
	}
	//v := <-noBufChan //invalid operation: <-noBufChan (receive from send-only type chan<- int)
	close(noBufChan)
	wt.Done()
}

func main() {

	wt.Add(2)

	//无缓冲的chan，同步方式，有读才能写入
	var noBufChan = make(chan int)
	//有缓冲的chan，一次性写入
	//var BufChan = make(chan int,5)

	go InChan(noBufChan)
	go OutChan(noBufChan)

	wt.Wait()

	fmt.Println("goroute全部退出")
}
