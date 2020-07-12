package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"myproj.lee/try/testrpc/protos"
)

const HelloServiceName = "path/to/pkg.HelloService"

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal("dialing:", err)
	}

	defer conn.Close()
	client := protos.NewHelloServiceClient(conn)
	a, _ := client.Hello(context.TODO(), &protos.String{Value: "kdjf"})

	fmt.Println(a.GetValue())
}
