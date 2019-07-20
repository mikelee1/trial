package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {

		go func() {

			fmt.Println(i)

			wg.Done()

		}()
	}
	wg.Wait()

}
