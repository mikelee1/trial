package main

import (
	"time"
	"fmt"
	"runtime"
)

func main()  {
	a := [100]int{}
	for i := 0; i <100; i++{
		go func(ii int) {
			for {
				a[ii]++
				runtime.Gosched()
			}

		}(i)
	}

	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
