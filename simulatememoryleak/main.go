package main

import (
	"time"
	"fmt"
)

func main() {
	for {
		time.Sleep(100*time.Microsecond)
		go func() {
			ticker := time.NewTicker(1*time.Second)
			select {
			case t := <-ticker.C:
				fmt.Println(t)
			}

		}()

	}
}
