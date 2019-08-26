package handler

import (
	"context"
	"errors"
	"github.com/op/go-logging"
	"github.com/openzipkin/zipkin-go"
	"google.golang.org/grpc/metadata"
	logger2 "myproj/try/common/logger"
	"myproj/try/testopenzipkin/protos"
	tracer2 "myproj/try/testopenzipkin/tracer"
)

var (
	logger *logging.Logger
	tracer *zipkin.Tracer
)

func init() {
	logger = logger2.GetLogger()
	tracer, _ = tracer2.NewTracer()
}

type HelloService struct{}

func (h HelloService) Hello(ctx context.Context, req *protos.String) (*protos.String, error) {
	//从grpc的client ctx中获取信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Info("not found md in context")
		return &protos.String{Value: "oops"}, errors.New("not found md in context")
	}
	logger.Info("md: ", md)
	//logger.Info(fmtstruct.String(md))
	//logger.Warning(fmtstruct.String(zipkin.SpanFromContext(ctx)))
	//上报
	span, ctx := tracer.StartSpanFromContext(ctx, "grpc server1")
	defer span.Finish()

	return &protos.String{Value: "hello," + req.Value}, nil
}
