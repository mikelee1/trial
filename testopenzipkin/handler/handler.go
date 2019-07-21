package handler

import (
	"github.com/op/go-logging"
	logger2 "myproj/try/common/logger"
	"context"
	"myproj/try/testopenzipkin/protos"
)

var logger *logging.Logger

func init()  {
	logger = logger2.GetLogger()
}

type HelloService struct {}

func (h HelloService)Hello(ctx context.Context, req *protos.String) (*protos.String, error)  {
	logger.Info("get request")
	return &protos.String{Value:"hello,"+req.Value},nil
}
