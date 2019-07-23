package logger

import (
	"github.com/op/go-logging"
	"os"
)

var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("common.logger")
	stdoutBackend := logging.NewBackendFormatter(
		logging.NewLogBackend(os.Stdout, "", 0),
		logging.MustStringFormatter(`%{color}[%{time:2006-01-02 15:04:05.000}] [%{module}] <%{shortfile}> %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`),
	)
	logging.SetBackend(stdoutBackend)
}

func GetLogger() *logging.Logger {
	if logger == nil {
		panic("no logger")
	}
	return logger
}
