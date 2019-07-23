package main

import (
	"aria/core"
	"fmt"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"os"
)

var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("test config")
}

type Config struct {
	Mode string
	Test
	Test1 string
}

type Test struct {
	Test1 string
}

func main() {
	var err error
	c := Config{}
	v := viper.New()
	cwd, err := os.Getwd()
	if err != nil {
		logger.Error(err)
		return
	}
	v.SetConfigFile(cwd + "/config/config.yaml")
	err = v.ReadInConfig()
	if err != nil {
		logger.Error(err)
		return
	}
	err = v.Unmarshal(&c)
	if err != nil {
		logger.Error(err)
		return
	}
	for _, v1 := range core.Flatten(c) {
		fmt.Println(v1)
	}
}
