package client

import (
	"context"
	"github.com/go-ego/ego/mid/json"
	"github.com/op/go-logging"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/propagation/b3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"math/rand"
	logger2 "myproj/try/common/logger"
	"myproj/try/common/ratelimit"
	"myproj/try/testopenzipkin/protos"
	tracer2 "myproj/try/testopenzipkin/tracer"
	"net/http"
	"time"
)

var (
	logger *logging.Logger
	tracer *zipkin.Tracer
	conn   *grpc.ClientConn
)

func init() {
	logger = logger2.GetLogger()
	tracer, _ = tracer2.NewTracer()
	conn, _ = grpc.Dial("127.0.0.1:1080", grpc.WithInsecure())
}

func CheckDB(ctx context.Context) {
	rand.Seed(time.Now().Unix())
	rs := rand.Intn(20) + 50
	logger.Infof("wait %d ms\n", rs)

	span, ctx := tracer.StartSpanFromContext(ctx, "CheckDB")
	defer span.Finish()
	//模拟延时
	time.Sleep(time.Duration(rs) * time.Millisecond)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	rl := ratelimit.NewLimit(3)
	defer rl.Release()
	logger.Info("start")
	defer logger.Info("complete")
	ctx := r.Context()
	//超时退出
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	//调用一次span
	CheckDB(ctx)

	//调用一次span
	span, ctx := tracer.StartSpanFromContext(ctx, "test")
	defer span.Finish()

	//调用一次span
	span, ctx = tracer.StartSpanFromContext(ctx, "test2")
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
	res, err := hsclient.Hello(ctx, &protos.String{Value: "mike"})
	if err != nil {
		logger.Error(err)
		return
	}

	//调用一次span
	span, ctx = tracer.StartSpanFromContext(ctx, "test1")
	defer span.Finish()

	resp, _ := json.Marshal(
		struct {
			StatusCode int
			Data       string
		}{http.StatusOK, res.Value},
	)
	w.Write(resp)
}
