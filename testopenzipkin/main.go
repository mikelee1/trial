package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/op/go-logging"
	"github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	"golang.org/x/net/netutil"
	_ "golang.org/x/net/netutil"
	"google.golang.org/grpc"
	logger2 "myproj/try/common/logger"
	"myproj/try/common/ratelimit"
	"myproj/try/testopenzipkin/client"
	"myproj/try/testopenzipkin/handler"
	"myproj/try/testopenzipkin/protos"
	tracer2 "myproj/try/testopenzipkin/tracer"
	"net"
)

var (
	err    error
	tracer *zipkin.Tracer
	logger *logging.Logger
)

func init() {
	logger = logger2.GetLogger()
}

func main() {
	tracer, err = tracer2.NewTracer()
	if err != nil {
		logger.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/hello", client.HelloHandler)

	//启动grpc的server端，内有限流器
	go StartGrpcServer()

	//中间件：对外server层的限流器
	rl := ratelimit.NewLimit(3)
	r.Use(func(f http.Handler) http.Handler {
		rl.Wait()
		return f
	})

	//中间件：设置server
	r.Use(zipkinhttp.NewServerMiddleware(
		tracer,
		zipkinhttp.SpanName("request")), // name for request span
	)

	err = http.ListenAndServe("0.0.0.0:18080", r)
	if err != nil {
		logger.Error(err)
		return
	}
}

func StartGrpcServer() {
	//生成grpcserver
	server := grpc.NewServer(grpc.StatsHandler(zipkingrpc.NewServerHandler(tracer)))

	//注册到上面的grpcserver
	protos.RegisterHelloServiceServer(server, handler.HelloService{})
	listener, err := net.Listen("tcp", ":1080")
	if err != nil {
		logger.Fatal("ListenTCP error:", err)
	}
	//grpc层的限流器
	listener = netutil.LimitListener(listener, 3)

	logger.Info("listening grpc request...")
	err = server.Serve(listener)
	if err != nil {
		logger.Fatal("fail to start grpc server")
	}
}
