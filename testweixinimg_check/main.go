package main

import (
	"fmt"
	"net/http"
	"github.com/op/go-logging"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"mime/multipart"
)

var logger = logging.MustGetLogger("test")

func main() {
	res := SendMsg()
	logger.Info(res)
}

type AccessToken struct {
	Access_token string
	Expires_in   int64
}

type InformRes struct {
	Errcode int64
	Errmsg  string
}

func SendMsg() bool {
	appid := "wx993a1ac03be0e1f4"
	secret := "28a41d4a8c4ec770dc628f16cd7b67bf"
	requestString := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?appid=%s&secret=%s&grant_type=client_credential", appid, secret)
	resp, err := http.Get(requestString)
	defer resp.Body.Close()
	if err != nil {
		logger.Info("sth wrong")
		logger.Info(err)
		return false
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var at AccessToken
	err = json.Unmarshal([]byte(string(body)), &at)
	if err != nil {
		logger.Info("cant unmarshal")
		logger.Error(err)
		return false
	}



	url := "https://api.weixin.qq.com/wxa/img_sec_check?access_token=" + at.Access_token
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	bin, err := ioutil.ReadFile("./testweixinimg_check/aaa.png")
	if err != nil {
		fmt.Println(err)
		return false
	}

	fw, err := w.CreateFormFile("file", "filename1")
	if err != nil {
		fmt.Println(err)
		return false
	}

	_, err = fw.Write(bin)
	if err != nil {
		fmt.Println(err)
		return false
	}
	w.Close()


	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		fmt.Println("req err: ", err)
		return false
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("resp err: ", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false
	}


	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info("fail to read all")
		logger.Info(err)
		return false
	}
	res := &InformRes{}
	err = json.Unmarshal(body, res)
	if err != nil {
		logger.Error(err)
		return false
	}
	if res.Errcode == 0 {
		logger.Info("nice")
		return true
	}
	logger.Info(res.Errmsg,res.Errcode)
	return false
}
