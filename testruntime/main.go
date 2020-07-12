package main

import (
	"fmt"
	"runtime"
	"math"
	"sync/atomic"
	"sync"
	"time"
)

var counter int64 = 0

func main() {
	start := time.Now()
	defer func() {
		atomic.LoadInt64(&counter)
		fmt.Println(time.Now().Sub(start).Seconds(), counter)
	}()
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(4)
	wg.Add(2)
	var e float64 = 7
	//go func() {
	//	for float64(counter) < math.Pow(10, e) {
	//		atomic.AddInt64(&counter, 1)
	//	}
	//	wg.Done()
	//}()
	//
	//go func() {
	//	for float64(counter) < math.Pow(10, e) {
	//		atomic.AddInt64(&counter, 1)
	//	}
	//	wg.Done()
	//}()
	//
	go func() {
		for float64(counter) < math.Pow(10, e) {
			atomic.AddInt64(&counter, 1)
		}
		wg.Done()
	}()

	go func() {
		for float64(counter) < math.Pow(10, e) {
			atomic.AddInt64(&counter, 1)
		}
		wg.Done()
	}()

	wg.Wait()
}
