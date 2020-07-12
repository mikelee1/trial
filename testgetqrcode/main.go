package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"github.com/op/go-logging"
	"os"
)

func main() {
	GetProblemQrcode("lee")
}

var logger logging.Logger

type AccessToken struct {
	Access_token string
	Expires_in   int64
}

type InformRes struct {
	Errcode int64
	Errmsg  string
}

func GetAccessToken() (AccessToken, bool) {
	var at AccessToken
	appid := "wx993a1ac03be0e1f4"
	secret := "28a41d4a8c4ec770dc628f16cd7b67bf"
	requestString := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?appid=%s&secret=%s&grant_type=client_credential", appid, secret)
	resp, err := http.Get(requestString)
	defer resp.Body.Close()
	if err != nil {
		logger.Error("sth wrong", err)
		return at, false
	}
	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal([]byte(string(body)), &at)
	if err != nil {
		logger.Error("cant unmarshal", err)
		return at, false
	}
	return at, true
}

func GetProblemQrcode(path string) (string, bool) {

	at, ok := GetAccessToken()
	if !ok {
		return "", false
	}
	url := "https://api.weixin.qq.com/wxa/getwxacode?access_token=" + at.Access_token
	body1 := map[string]interface{}{
		"path": path,
	}
	bs, _ := json.Marshal(body1)
	req := bytes.NewBuffer([]byte(bs))

	body_type := "application/json;charset=utf-8"

	resp, err := http.Post(url, body_type, req)
	defer resp.Body.Close()
	if err != nil {
		logger.Error("sth wrong", err)
		return "", false
	}
	body, _ := ioutil.ReadAll(resp.Body)
	bodyraw := body
	//
	res := &InformRes{}
	err = json.Unmarshal(body, res)
	if err != nil {
		f, err := os.OpenFile("testgetqrcode/lee.jpg", os.O_CREATE|os.O_RDWR, os.ModePerm)
		defer f.Close()
		if err != nil {
			logger.Error("sth wrong", err)
			return "", false
		}
		_, err = f.Write(bodyraw)
		if err != nil {
			logger.Error("sth wrong", err)
			return "", false
		}
		return "", true
	}
	return "", false
}
