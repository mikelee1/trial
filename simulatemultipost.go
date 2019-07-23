package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var wg11 *sync.WaitGroup

func main() {
	timer1 := time.NewTicker(1 * time.Second)
	defer timer1.Stop()
	go func() {
		for {
			select {
			case <-timer1.C:
				fmt.Println(time.Now())
			}
		}

	}()
	wg11 = &sync.WaitGroup{}
	for range [3]int{} {
		wg11.Add(1)

		go Post("http://192.168.9.87:8081/indirect/invoke")
		time.Sleep(10 * time.Second)
	}
	wg11.Wait()
}

type Request struct {
	Invoke      bool
	Channelname string
	Ccname      string
	Args        []string
	Userid      string
}

func Post(url string) string {
	defer func() {
		wg11.Done()
	}()
	Tmpa := &Request{
		Invoke:      true,
		Channelname: "channel1",
		Ccname:      "example1",
		Args:        []string{"move", "a", "b", "3"},
		Userid:      "21470184-21bc-4996-8f47-474f5dfba25b",
	}
	jsonValue, _ := json.Marshal(Tmpa)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		panic(err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return string(body)
}
