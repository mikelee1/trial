package main

import (
	"fmt"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"myproj.lee/try/common/flatteninterface"
	logger2 "myproj.lee/try/common/logger"
	"os"
)

var logger *logging.Logger

func init() {
	logger = logger2.GetLogger()
}

type Config struct {
	Mode    string
	Test
	Test1   string
	Clients []Client
}

type Client struct {
	Name string
	Ip   string
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
	for _, v1 := range flatteninterface.Flatten(c) {
		fmt.Println(v1)
	}
}
