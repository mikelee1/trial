package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {

	go func() {
		var i = 0
		for {
			go getten()
			i++
			if i > 10 {
				break
			}
		}

	}()

	time.Sleep(10 * time.Second)

}

func getten() {
	for {
		time.Sleep(100 * time.Millisecond)
		resp, err := http.Post("https://seucareer.liyuanye.club/topic/getten",
			"application/json;charset=utf-8", strings.NewReader(`{"userid":1,"topictype":"","pagenum":0,"pagesize":10}`))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("getten")
	}

}
