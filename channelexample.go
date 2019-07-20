package main

import (
	"fmt"
	"time"
)

var lis = [10]int{1,2,3}
func consume1(a <-chan int,down chan bool)  {//只能读

	for {
		select {
		case i:=<-a:
			//time.Sleep(time.Second*1)
			fmt.Printf("consume1:%v\n",i)

		}
	}
}


func consume2(a <-chan int,down chan bool)  {//只能读

	for {
		select {
		case i:=<-a:
			time.Sleep(time.Second*1)
			fmt.Printf("consume2:%v\n",i)

		}
	}
}

func produce(a chan<- int,down chan bool)  {//只能写
	for i := range lis{
		//time.Sleep(time.Second*1)
		a <- i
		fmt.Printf("produce:%v\n",i)
	}
	time.Sleep(time.Second*2)
	down <- true
}


func main()  {
	var a = make(chan int,10)//size要设置对
	var down = make(chan bool,1)

	go produce(a,down)
	time.Sleep(time.Second)
	go consume1(a,down)
	go consume2(a,down)

	if <-down{
		fmt.Print("down\n")
	}

}
