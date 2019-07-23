package main_test

import (
	"testing"
	logger2 "myproj/try/common/logger"
	"github.com/op/go-logging"
	"google.golang.org/grpc"
	"myproj/try/testopenzipkin/protos"
	"context"
	"sync"
	"time"
	"net/http"
	"fmt"
)

var logger *logging.Logger
var wg *sync.WaitGroup

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

//测试限流器
func Test_main1(t *testing.T) {
	wg = &sync.WaitGroup{}
	wg.Add(4)
	go OneTime()
	go OneTime()
	go OneTime()
	go OneTime()
	wg.Wait()
}

func OneTime()  {
	resp,err := http.Get("http://127.0.0.1:8080/hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
	fmt.Println(time.Now().Unix())
	wg.Done()
}