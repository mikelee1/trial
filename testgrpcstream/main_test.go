package main_test

import (
	"testing"
	"google.golang.org/grpc"
	"myproj/try/testgrpcstream/user"
	"context"
	"fmt"
	"time"
)

const (
	ADDRESS = "localhost:4000"
)

func Test_main(t *testing.T) {
	//通过grpc 库 建立一个连接
	conn ,err := grpc.Dial(ADDRESS,grpc.WithInsecure())
	if err != nil{
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := testuser.NewEchoServiceClient(conn)

	//服务端 客户端 双向流
	allStr,_ := c.Echo(context.Background())
	c1 := make(chan int32,1)
	go func() {
		for {
			data,_ := allStr.Recv()
			if data != nil{
				c1 <- data.B
				fmt.Println(data)
			}
		}
	}()

	go func() {
		for {
			select {
			case bdata := <- c1:
				allStr.Send(&testuser.EchoRequest{A:bdata+1})
				time.Sleep(time.Second)
			}
		}
	}()
	c1 <- 1
	select {
	}

}

func Test_Pingpong(t *testing.T) {
	conn,_ := grpc.Dial(ADDRESS,grpc.WithInsecure())
	ppclient := testuser.NewEchoServiceClient(conn)
	stream,_ := ppclient.PingPong(context.TODO())
	c1 := make(chan string,1)
	go func() {
		for {
			data,_ := stream.Recv()
			if data != nil{
				c1 <- data.Data
				fmt.Println(data)
			}
		}
	}()

	go func() {
		for {
			select {
			case bdata := <- c1:
				stream.Send(&testuser.PingRequest{Data:bdata})
				time.Sleep(time.Second)
			}
		}
	}()

	c1 <- "ping"
	select {

	}
}