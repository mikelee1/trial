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
	"google.golang.org/grpc/metadata"
	"github.com/openzipkin/zipkin-go/propagation/b3"
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
	span,ctx := tracer.StartSpanFromContext(ctx,"CheckDB")
	defer span.Finish()
}

func HelloHandler(w http.ResponseWriter, r *http.Request)  {
	ctx := r.Context()

	//调用一次span
	CheckDB(ctx)

	//调用一次span
	span,ctx := tracer.StartSpanFromContext(ctx,"test")
	defer span.Finish()

	//调用一次span
	span,ctx = tracer.StartSpanFromContext(ctx,"test2")
	defer span.Finish()

	//向grpc中存储父span
	md := &metadata.MD{"key1": []string{"val1"}, "key2": []string{"val2"}}
	err := b3.InjectGRPC(md)(span.Context())
	if err != nil {
		logger.Error(err)
		return
	}
	ctx = metadata.NewOutgoingContext(ctx, *md)

	hsclient := protos.NewHelloServiceClient(conn)
	res,err := hsclient.Hello(ctx,&protos.String{Value:"mike"})
	if err != nil {
		logger.Error(err)
		return
	}

	//调用一次span
	span,ctx = tracer.StartSpanFromContext(ctx,"test1")
	defer span.Finish()

	resp,_ := json.Marshal(
		struct {
			StatusCode int
			Data string
		}{http.StatusOK,res.Value},
	)
	w.Write(resp)
}
