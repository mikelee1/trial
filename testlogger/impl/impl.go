package impl

import (
	"fmt"
	"github.com/op/go-logging"
)

var logger *logging.Logger

func init() {

	logger = logging.MustGetLogger("config")
	fmt.Println("in impl")
}

func Run() {
	logger.Infof("what")
	logger.Error("err")
	logger.Fatal("fatal")
	logger.Info("info")
}
