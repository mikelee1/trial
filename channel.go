package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.
	messages := make(chan string, 1)
	res := make(chan string, 1)
	clo1 := make(chan string, 1)
	clo2 := make(chan string, 1)

	// _Send_ a value into a channel using the `channel <-`
	// syntax. Here we send `"ping"`  to the `messages`
	// channel we made above, from a new goroutine.
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {

		defer wg.Done()
		//time.Sleep(2*time.Second)
		messages <- "ping"
		sign := false
		for {
			select {
			case b := <-res:
				fmt.Println(b)
				time.Sleep(1 * time.Second)
				messages <- "ping" + b
			case <-clo1:
				sign = true
				break
			case <-time.After(time.Second * 2):
				fmt.Println("timeout 2 seconds")
			}
			if sign {
				break
			}

		}

		fmt.Println("doen func1")

	}(&wg)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {

		defer wg.Done()
		sign := false
		for {
			select {
			case a := <-messages:
				fmt.Println(a)
				time.Sleep(1 * time.Second)
				res <- a
			case <-clo2:
				sign = true
				break
			case <-time.After(time.Second * 2):
				fmt.Println("timeout 2 seconds")
			}
			if sign {
				break
			}

		}
		fmt.Println("doen func2")

	}(&wg)
	// The `<-channel` syntax _receives_ a value from the
	// channel. Here we'll receive the `"ping"` message
	// we sent above and print it out.
	//msg := <-messages
	//fmt.Println(msg)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {

		defer wg.Done()

		time.Sleep(10 * time.Second)
		clo1 <- "close"
		clo2 <- "close"

	}(&wg)

	wg.Wait()
}
