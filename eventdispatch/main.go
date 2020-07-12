package main

import (
	"github.com/op/go-logging"
	e "myproj.lee/try/eventdispatch/core"
	"os"
	"time"
)

const HELLO_WORLD = "helloWorld"
const Mike = "mike"
const Sample = "sample"

var logger *logging.Logger

func init() {
	stdoutBackend := logging.NewBackendFormatter(
		logging.NewLogBackend(os.Stdout, "", 0),
		logging.MustStringFormatter(`%{color}[%{time:2006-01-02 15:04:05.000}] [%{module}] <%{shortfile}> %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`),
	)
	logging.SetBackend(stdoutBackend)

	logger = logging.MustGetLogger("main")
}

func main() {
	//mike 新建handler
	dispatcher := e.NewEventDispatcher()
	//mike 绑定处理函数
	listener := e.NewEventListener(myEventListener)
	//mike 相当于注册事件到listener上
	dispatcher.AddEventListener(HELLO_WORLD, listener)
	dispatcher.AddEventListener(Mike, listener)

	time.Sleep(time.Second * 2)
	logger.Info("to remove")
	res := dispatcher.RemoveEventListener(HELLO_WORLD, listener)
	if res {
		logger.Info("removed")
	} else {
		logger.Info("no removed")
	}
	//mike 分发新建的event
	dispatcher.DispatchEvent(e.NewEvent(HELLO_WORLD, nil))
	res = dispatcher.DispatchEvent(e.NewEvent(Sample, nil))
	if !res {
		logger.Error("e")
	}
}

//mike 处理函数
func myEventListener(event e.Event) {

	logger.Info(event.Type, event.Object, event.Target)
}
