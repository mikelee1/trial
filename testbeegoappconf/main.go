package main

import (
	"github.com/astaxie/beego"
	"fmt"
	logger2 "myproj/try/common/logger"
	"time"
	"path/filepath"
)


func main() {
	var err error
	logger := logger2.GetLogger()
	path := filepath.Join("testbeegoappconf", "conf", "app.conf")
	logger.Info(path)
	err = beego.LoadAppConfig("ini",path)
	if err != nil {
		logger.Error(err)
		return
	}
	appconfig := beego.AppConfig
	logger.Info(appconfig.String("appname"))

	logger.Infof("%#v",appconfig)
	time.Sleep(time.Second)
	err = appconfig.Set("appname","aaa")
	if err != nil {
		logger.Error(err)
		return
	}
	a := appconfig.String("a")
	err = appconfig.SaveConfigFile("testbeegoappconf/conf/testbeego.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)
}
