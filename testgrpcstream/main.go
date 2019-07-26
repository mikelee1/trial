package main

import (
	"myproj/try/testgrpcstream/user"
	"net"
	"google.golang.org/grpc"
	"fmt"
	"sync"
	"time"
)

func main() {
	lis,err := net.Listen("tcp",":4000")
	if err != nil{
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()
	//注册事件
	testuser.RegisterEchoServiceServer(s,&UserEcho{})
	//处理链接
	fmt.Println("start listen...")
	err = s.Serve(lis)
	if err != nil {
		fmt.Println(err)
		return
	}
}

type UserEcho struct {}

func (u *UserEcho)PingPong(p testuser.EchoService_PingPongServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			r,_ := p.Recv()
			if r != nil{
				fmt.Println(r)
			}
		}
	}()

	go func() {
		for {
			p.Send(&testuser.PongResponse{Data:"pong"})
			time.Sleep(time.Second)
		}
	}()



	select {

	}
	return nil
}



func (u *UserEcho)Echo(a testuser.EchoService_EchoServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	c := make(chan int32,1)
	go func() {
		for {
			data, _ := a.Recv()
			if data != nil{
				fmt.Println(data)
				c <- data.A
			}
		}
		wg.Done()
	}()

	go func() {
		for {
			select {
			case adata := <-c:
				a.Send(&testuser.EchoResponse{B: adata+1})
				time.Sleep(2*time.Second)
			}
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}