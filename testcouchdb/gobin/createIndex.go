package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-vgo/gt/file"
	"io/ioutil"
	"net/http"
	"flag"
	"errors"
)

//本地测试
//./main -username=admin -password=pass -couchdbip=192.168.9.87:31994 -channelcc=mychannel_ccdc -indexpath=./META-INF
//内网测试
//./createIndex -username=admin -password=yunphant -couchdbip=192.168.9.164:31994 -channelcc=testchannel_testccdc -indexpath=./META-INF

var (
	Cookie    string
	Username  string
	Password  string
	CouchdbIp string //couchdb的服务地址
	Channelcc string //通道名_合约名
	IndexPath string //index的目录
)

func init() {
	flag.StringVar(&Username, "username", "", "输入couchdb用户名")
	flag.StringVar(&Password, "password", "", "输入couchdb密码")
	flag.StringVar(&CouchdbIp, "couchdbip", "192.168.9.87:31994", "输入couchdb的ip:port")
	flag.StringVar(&Channelcc, "channelcc", "mychannel_ccdc", "输入通道名_合约名")
	flag.StringVar(&IndexPath, "indexpath", "./testcouchdb/gobin/META-INF", "输入index的目录")
}

type ReqLogin struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

func Login() error {
	l := ReqLogin{
		Name:     Username,
		Password: Password,
	}
	lb, err := json.Marshal(l)
	if err != nil {
		return err
	}
	resp, err := http.Post(fmt.Sprintf("http://%s/_session", CouchdbIp), "application/json", bytes.NewBuffer(lb))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	cookies := resp.Cookies()
	Cookie = fmt.Sprintf("%s=%s", cookies[0].Name, cookies[0].Value)
	return nil
}

type RespIndex struct {
	Result string `json:"result"`
	Id     string `json:"id"`
	Name   string `json:"name"`

	Error  string `json:"error"`
	Reason string `json:"reason"`
}

func CreateIndex(a []byte, fileName string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/%s/_index", CouchdbIp, Channelcc), bytes.NewBuffer(a))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Cookie", Cookie)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	res := RespIndex{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	if res.Result != "created" {
		fmt.Printf("[ %s ] fileName %s\n", res.Result, fileName)
		return nil
	}
	fmt.Printf("[ ok ]  fileName %s\n", fileName)
	return nil
}

func WalkGoFile(dir string) ([]string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic error: ", err)
		}
	}()
	a, err := file.Walk(dir, "")
	if err != nil {
		return []string{}, err
	}
	return a, nil
}

func main() {
	var err error
	flag.Parse()
	//fmt.Println(Username, Password, CouchdbIp, Channelcc, IndexPath)
	if Username != "" || Password != "" {
		err = Login()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	files, err := WalkGoFile(IndexPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range files {
		jsonBytes, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = CreateIndex(jsonBytes, f)
		if err != nil {
			fmt.Println("err :", err)
			return
		}
	}
}
