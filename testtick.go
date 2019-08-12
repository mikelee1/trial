package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.NewTicker(2*time.Second)
	defer ticker.Stop()
	for {
		//time.Sleep(time.Second)
		select {
		case t := <-ticker.C:
			fmt.Println(t)
		//default:
		//	fmt.Println("default")

		}
	}
}
