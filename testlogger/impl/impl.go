package impl

import (
	"github.com/op/go-logging"
	"fmt"
)

var logger *logging.Logger

func init() {

	logger = logging.MustGetLogger("config")
	fmt.Println("in impl")
}

func Run()  {
	logger.Infof("what")
}
