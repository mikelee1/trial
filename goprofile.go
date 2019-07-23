package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("hello world")
		}
	}()
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
