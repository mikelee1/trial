package main

import (
	"io/ioutil"
	"github.com/op/go-logging"
	logger2 "myproj/try/common/logger"
	"os"
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"github.com/magiconair/properties/assert"
	"testing"
)

var logger *logging.Logger

func init()  {
	logger = logger2.GetLogger()
}

func main()  {
	t := &testing.T{}
	//文件内容一样，hash的结果就一样
	a := Hash("hashfile/a.log")
	b := Hash("hashfile/b.log")
	assert.Equal(t,a,b)
}

func Hash(filename string) string {
	dchaincode,err := os.Open(filename)
	if err != nil {
		logger.Error(err)
		return ""
	}
	chaincodecontent, err := ioutil.ReadAll(dchaincode)
	if err != nil {
		logger.Error(err)
		return ""
	}
	hash := CalFileMD5(chaincodecontent)
	logger.Info(hash)
	return hash
}

func CalFileMD5(content []byte) string {
	//生成MD5 HASH摘要
	reader := bytes.NewReader(content)
	md5h := md5.New()
	io.Copy(md5h, reader)
	ans := md5h.Sum([]byte(""))
	ansstr := fmt.Sprintf("%x", ans)
	return ansstr
}