package main

import (
	"fmt"
	"time"
)

func main() {
	goroutine()
	normal()
}

//用生产者消费者模式，时长按照生产者和消费者中耗时久的来
func goroutine() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Now().Sub(start).Seconds())
	}()
	dataC := make(chan int)
	exitC := make(chan bool)

	go func() {
		for d := range dataC {
			fmt.Println(d)
			time.Sleep(time.Second * 2)
		}
		exitC <- true
	}()

	for _, d := range []int{1, 2, 3} {
		dataC <- d
		time.Sleep(time.Second)
	}
	close(dataC)
	<-exitC
}

//不用生产者消费者模式，时长增加一倍
func normal() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Now().Sub(start).Seconds())
	}()
	data := []int{}
	for _, d := range []int{1, 2, 3} {
		data = append(data, d)
		time.Sleep(time.Second)
	}

	for d := range data {
		fmt.Println(d)
		time.Sleep(time.Second)
	}
}
