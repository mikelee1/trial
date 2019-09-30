package main

import (
	"os"
	logger2 "myproj/try/common/logger"
	"github.com/op/go-logging"
	"io/ioutil"
	"strings"
)

var logger *logging.Logger

func init() {
	logger = logger2.GetLogger()
}

func main() {
	//获取当前工程路径
	dir, _ := os.Getwd()
	logger.Info(dir)
	//遍历当前目录下的文件
	listFile(dir + "/testosdir")
}

func listFile(myfolder string) {
	files, _ := ioutil.ReadDir(myfolder)
	for _, file := range files {
		if file.IsDir() {
			listFile(myfolder + "/" + file.Name())
		} else {
			if strings.HasSuffix(file.Name(), "jpg") {
				logger.Info("got jpg")
				err := os.Rename(myfolder+"/"+file.Name(), myfolder+"/"+strings.Replace(file.Name(), ".jpg", "tmp.jpg", -1))
				if err != nil {
					logger.Error(err)
					return
				}
			}
			logger.Info(myfolder + "/" + file.Name())
		}
	}
}
