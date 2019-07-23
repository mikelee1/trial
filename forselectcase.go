package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	stopch := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 1
		ch1 <- 2
		time.Sleep(1 * time.Second)
		stopch <- true

	}()
	go func() {
		for {
			fmt.Println("in for")
			time.Sleep(500 * time.Millisecond)
			select {
			case d1 := <-ch1:
				fmt.Println("ch1", d1)

			case d2 := <-ch2:
				fmt.Println("ch2", d2)
			}
		}
	}()
	<-stopch
	close(ch1)
	close(ch2)
	close(stopch)

}
