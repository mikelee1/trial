package pkg

import (
	"fmt"
	"time"
)

func init() {
	go func() {
		for true {
			fmt.Println("11")
			time.Sleep(time.Second)
		}
	}()
}
