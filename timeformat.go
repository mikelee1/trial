package main

import (
	"time"
	"fmt"
	"github.com/op/go-logging"
)
var format = logging.MustStringFormatter(
	`%{color}%{time:2006-01-02 15:04:05.000 MST} [%{shortfile}] [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x} %{message}%{color:reset}`,
)
func main()  {
	//var a = time.Now().Format("2017-08-04")
	//
	//fmt.Printf("today is:%v\n",a)
	//fmt.Printf("now is:%v\n",time.Now().Unix())
	//
	//
	//b := time.Now().Format("2006-01-02")
	////fmt.Printf("b is:%s\n",b)
	//
	//logfile,_ := os.OpenFile("logs/"+b+".txt",os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	//
	//log,_ := logging.GetLogger("test")
	//log1,_ := logging.GetLogger("yes1")
	//
	//backend := logging.NewLogBackend(logfile,"",1)
	//backendformat := logging.NewBackendFormatter(backend,format)
	//logging.SetBackend(backendformat).SetLevel(logging.DEBUG,"what")
	//backendleveled := logging.AddModuleLevel(backendformat)
	//backendleveled.SetLevel(logging.DEBUG,"what")
	//log.Info(b)
	//log1.Info("yes")
	//
	//log.Info(time.Now().Day())
	//
	//
	//go BoottimeTimingSettlement()
	////var ch chan interface{}
	////go demo(ch)
	//time.Sleep(time.Minute*1)

	//Test_1()


	time := time.Now()
	// 默认UTC

	fmt.Println(time.Local())
	// 一般为CST
	//loc, _ := time.LoadLocation("Local")
	//
	//// CST
	//loc, _:= time.Local("Asia/Chongqing")


}

func BoottimeTimingSettlement() {
	now := time.Now()
	next := now.Add(time.Second * 10)
	next = time.Date(next.Year(), next.Month(), now.Day(), now.Hour(), now.Minute(), next.Second(), 0, next.Location())
	t := time.NewTimer(next.Sub(now))
	for {
		select {
			case <-t.C:
			fmt.Printf("完成: %v\n",time.Now())
			t.Reset(time.Second*10)//10秒后重新t.c
		}
	}
}



func demo(input chan interface{}) {
	t1 := time.NewTimer(time.Second * 5)
	t2 := time.NewTimer(time.Second * 10)

	for {
		select {
		case msg := <-input:
			println(msg)

		case <-t1.C:
			println("5s timer")
			t1.Reset(time.Second * 5)

		case <-t2.C:
			println("10s timer")
			t2.Reset(time.Second * 10)
		}
	}
}

func Test_1()  {
	t := time.Now()
	fmt.Printf(".......to processproposal:%v\n",t)
	fmt.Printf("---%v",time.Since(t))
}