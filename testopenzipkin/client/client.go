package client

import (
	"net/http"
	"google.golang.org/grpc"
	"myproj/try/testopenzipkin/protos"
	"context"
	"github.com/op/go-logging"
	logger2 "myproj/try/common/logger"
	"github.com/go-ego/ego/mid/json"
	"github.com/openzipkin/zipkin-go"
	tracer2 "myproj/try/testopenzipkin/tracer"
)

var (
	logger *logging.Logger
	tracer *zipkin.Tracer
	conn *grpc.ClientConn
)

func init()  {
	logger = logger2.GetLogger()
	tracer, _ = tracer2.NewTracer()
	conn,_ = grpc.Dial("127.0.0.1:1080",grpc.WithInsecure())
}

func CheckDB(ctx context.Context)  {
	span,ctx := tracer.StartSpanFromContext(ctx,"CheckDB",[]zipkin.SpanOption{}...)
	defer span.Finish()
}

func HelloHandler(w http.ResponseWriter, r *http.Request)  {
	ctx := r.Context()
	hsclient := protos.NewHelloServiceClient(conn)
	res,err := hsclient.Hello(ctx,&protos.String{Value:"mike"})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("res is: ",res)
	if err != nil {
		logger.Error(err)
		return
	}

	CheckDB(ctx)

	span,ctx := tracer.StartSpanFromContext(ctx,"test",[]zipkin.SpanOption{}...)
	defer span.Finish()

	span,ctx = tracer.StartSpanFromContext(ctx,"test1",[]zipkin.SpanOption{}...)
	defer span.Finish()

	resp,_ := json.Marshal(
		struct {
			StatusCode int
			Data string
		}{http.StatusOK,res.Value},
	)
	w.Write(resp)
}
