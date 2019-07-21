package main_test

import (
	"testing"
	logger2 "myproj/try/common/logger"
	"github.com/op/go-logging"
	"google.golang.org/grpc"
	"myproj/try/testopenzipkin/protos"
	"context"
)

var logger *logging.Logger

func init()  {
	logger = logger2.GetLogger()
}

func Test_main(t *testing.T) {
	conn,err := grpc.Dial("127.0.0.1:1080",grpc.WithInsecure())
	if err != nil {
		logger.Error(err)
		return
	}
	hsclient := protos.NewHelloServiceClient(conn)
	logger.Info(hsclient)
	res,err := hsclient.Hello(context.Background(),&protos.String{Value:"mike"})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("res is: ",res)
	if err != nil {
		logger.Error(err)
	}
}
