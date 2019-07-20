package main

import (
	"net/rpc"
	"net"
	"log"
	"fmt"
	"myproj/try/testrpc/protos"
	"google.golang.org/grpc"
	"context"
)

//type HelloService struct {}
//
//func (p *HelloService) Hello(request string, reply *string) error {
//	*reply = "hello:" + request
//	return nil
//}

type HelloService1 struct {}

func (p *HelloService1) Hello(ctx context.Context,args *protos.String) (*protos.String,error) {
	return  &protos.String{Value:"hello1:" + args.GetValue()},nil
}

func main() {
	grpcServer := grpc.NewServer()
	protos.RegisterHelloServiceServer(grpcServer,new(HelloService1))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		fmt.Println("create listener")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		fmt.Println("end listener")
		go rpc.ServeConn(conn)
	}
}


const HelloServiceName = "path/to/pkg.HelloService"

//type HelloServiceInterface = interface {
//	Hello(request protos.String, reply *string) error
//}

//func RegisterHelloService(svc HelloServiceInterface) error {
//	return rpc.RegisterName(HelloServiceName, svc)
//}