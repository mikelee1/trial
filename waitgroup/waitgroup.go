package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	li := make([]int, 5)
	for i := range li {
		wg.Add(1)
		go func(i int, group *sync.WaitGroup) {
			fmt.Println(i)
			wg.Done()
		}(i, wg)
	}
	wg.Wait()
}
