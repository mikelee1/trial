package main

import (
	"fmt"
	"github.com/op/go-logging"
	"myproj/try/testlogger/impl"
	"os"
)

func init() {
	stdoutBackend := logging.NewBackendFormatter(
		logging.NewLogBackend(os.Stdout, "", 0),
		logging.MustStringFormatter(`%{color}[%{time:2006-01-02 15:04:05.000}] [%{module}] <%{shortfile}> %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`),
	)
	logging.SetBackend(stdoutBackend)
	fmt.Println("in main init")
}

func main() {
	impl.Run()
}
