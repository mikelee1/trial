package main_test

import (
	"testing"
	"net/http"
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup

func Test_main(t *testing.T) {
	wg = &sync.WaitGroup{}
	wg.Add(2)
	go OneTime()
	time.Sleep(2000*time.Millisecond)
	go OneTime()
	wg.Wait()
}

func OneTime()  {
	resp,err := http.Get("http://127.0.0.1:8080/hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
	fmt.Println(time.Now().Unix())
	wg.Done()
}
