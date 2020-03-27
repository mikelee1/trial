package main_test

import (
	"testing"
	"time"
	"fmt"
	"sync"
)

//channel多路复用
func Test_MultiChannel(t *testing.T) {
	var wg sync.WaitGroup
	start := time.Now()
	defer func() {
		fmt.Println(time.Now().Sub(start).Seconds())
	}()
	highC := make(chan int)
	lowC := make(chan int)
	wg.Add(3)
	go func() {
		for {
			select {
			case d, ok := <-highC:
				if !ok {
					highC = nil
					break
				}
				fmt.Println("highC: ", d)
			case b, ok := <-lowC:
				if !ok {
					lowC = nil
					break
				}
				fmt.Println("lowC: ", b)
			}
			if highC == nil && lowC == nil {
				break
			}
		}
		wg.Done()
	}()

	go func() {
		for _, d := range []int{1, 2, 3} {
			highC <- d
			time.Sleep(time.Second)
		}
		close(highC)
		wg.Done()
	}()

	go func() {
		for _, d := range []int{1, 2, 3, 4, 5} {
			lowC <- d
			time.Sleep(time.Second)
		}
		close(lowC)
		wg.Done()
	}()

	wg.Wait()
}

//有优先级的channel
func Test_PriorityChannel(t *testing.T) {
	var wg sync.WaitGroup
	start := time.Now()
	defer func() {
		fmt.Println(time.Now().Sub(start).Seconds())
	}()
	highC := make(chan int)
	lowC := make(chan int)
	wg.Add(3)
	go func() {
		for {
			select {
			case data, ok := <-highC:
				if !ok {
					highC = nil
					break
				}
				handleHigh(data)
			default:
				select {
				case data, ok := <-highC:
					if !ok {
						highC = nil
						break
					}
					handleHigh(data)
				case data, ok := <-lowC:
					if !ok {
						lowC = nil
						break
					}
					handleLow(data)
				}
			}
			if highC == nil && lowC == nil {
				break
			}
		}

		wg.Done()
	}()

	go func() {
		for _, d := range []int{1, 2, 3} {
			highC <- d
			time.Sleep(time.Second)
		}
		close(highC)
		wg.Done()
	}()

	go func() {
		for _, d := range []int{1, 2, 3, 4, 5} {
			lowC <- d
			time.Sleep(time.Second)
		}
		close(lowC)
		wg.Done()
	}()

	wg.Wait()
}

func handleHigh(data interface{}) {
	fmt.Println(data)
}

func handleLow(data interface{}) {
	fmt.Println(data)
}
