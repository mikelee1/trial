package main

import (
	"net/http"

	"github.com/gorilla/mux"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	tracer2 "myproj/try/testopenzipkin/tracer"
	"github.com/openzipkin/zipkin-go"
	"myproj/try/testopenzipkin/protos"
	"google.golang.org/grpc"
	"net"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	"github.com/op/go-logging"
	logger2 "myproj/try/common/logger"
	"myproj/try/testopenzipkin/handler"
	"myproj/try/testopenzipkin/client"
)

var (
	err error
	tracer *zipkin.Tracer
	logger *logging.Logger
)

func init()  {
	logger = logger2.GetLogger()
}

func main() {
	tracer, err = tracer2.NewTracer()
	if err != nil {
		logger.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/hello", client.HelloHandler)

	//启动grpc的server端
	go StartGrpcServer()

	//设置server的中间件
	r.Use(zipkinhttp.NewServerMiddleware(
		tracer,
		zipkinhttp.SpanName("request")), // name for request span
	)

	http.ListenAndServe("0.0.0.0:8080",r)
}

func StartGrpcServer()  {
	//生成grpcserver
	server := grpc.NewServer(grpc.StatsHandler(zipkingrpc.NewServerHandler(tracer)))

	//注册到上面的grpcserver
	protos.RegisterHelloServiceServer(server,handler.HelloService{})
	listener, err := net.Listen("tcp", ":1080")
	if err != nil {
		logger.Fatal("ListenTCP error:", err)
	}

	logger.Info("listening grpc request...")
	err = server.Serve(listener)
	if err != nil{
		logger.Fatal("fail to start grpc server")
	}
}
