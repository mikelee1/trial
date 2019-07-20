package service

import (
	"github.com/hyperledger/fabric/sdk"
	"os"
	"errors"
	"github.com/op/go-logging"
)

var logger *logging.Logger

func init()  {
	logger = logging.MustGetLogger("service")
}

func GetCA(dir string, msp string) (*sdk.CA, error) {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return sdk.NewCA(dir, msp)
	}
	if err == nil {
		if !info.IsDir() {
			return nil, errors.New("msp path is not a directory, but a file")
		}
		return sdk.ConstructCAFromDir(dir)
	}
	return nil, err
}
