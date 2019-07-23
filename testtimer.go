package main

import (
	"fmt"
	"time"
)

//func main() {
//	t := time.NewTimer(5*time.Second)
//	defer t.Stop()
//	go func() {
//		<-t.C
//		fmt.Println("running 5s")
//	}()
//	time.Sleep(6*time.Second)
//}

func main() {
	t := time.NewTimer(5 * time.Second)
	defer t.Stop()
	c := make(chan int, 1)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

	}()

	for {
		t.Reset(1 * time.Second)
		select {
		case a := <-c:
			time.Sleep(100 * time.Millisecond)
			fmt.Println("a:", a)
		case <-t.C:
			fmt.Println("timeout")
			close(c)
			return
		}
	}

	//time.Sleep(6*time.Second)
}
