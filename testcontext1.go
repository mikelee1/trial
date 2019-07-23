package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
	"math/rand"
	"sync"
)

var (
	wg1 sync.WaitGroup
)

func work(ctx context.Context) error {
	defer wg1.Done()
	go work1(ctx)
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		randtime := rand.Int63n(8) + 1
		fmt.Println("randtime:", randtime)
		select {
		case <-time.After(time.Duration(randtime) * time.Second):
			fmt.Println("Doing some work ", i)
			//fmt.Println(ctx.Deadline())

			// we received the signal of cancelation in this channel
		case <-ctx.Done():
			fmt.Println("Cancel the context ", i)
			return ctx.Err()
		}
	}
	return nil
}

func work1(ctx context.Context) error {
	defer wg1.Done()
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		randtime := rand.Int63n(8) + 1
		fmt.Println("work1 randtime:", randtime)
		select {
		case <-time.After(time.Duration(randtime) * time.Second):
			fmt.Println("work1 Doing some work ", i)
		// we received the signal of cancelation in this channel
		case <-ctx.Done():
			fmt.Println("work1 Cancel the context ", i)
			return ctx.Err()
		}
	}
	return nil
}

func main() {
	//9秒后ctx.done()触发
	//ctx, cancel := context.WithTimeout(context.Background(), 9*time.Second)
	//10秒后ctx.done()幂等触发
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()

	fmt.Println("Hey, I'm going to do some work")

	wg1.Add(2)
	go work(ctx)
	//go work(ctx)
	wg1.Wait()

	fmt.Println("Finished. I'm going home")
}
