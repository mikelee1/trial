package main

import (
	"fmt"
	"net/http"
	"github.com/op/go-logging"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

var logger = logging.MustGetLogger("test")

func main() {
	openid := "oowmZ5R311ND3StOd4KBOUiT-XJI"
	formid := "2b144fc6f3a14833a4b3b3daff1accf3"
	res := SendInform(openid, formid, "ZmMHIF5bXgOSn_nspvWqPlpPxDHvmNyNRNytT2rkopY", "1", map[string]map[string]string{})
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

func SendInform(openid, formid, template_id, page string, data map[string]map[string]string) bool {
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

	url := "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=" + at.Access_token
	body1 := map[string]interface{}{
		"type":        "text",
		"touser":      openid,
		"template_id": template_id,
		"form_id":     formid,
		"data":        data,
		"page":        page,
	}
	bs, _ := json.Marshal(body1)
	req := bytes.NewBuffer([]byte(bs))

	body_type := "application/json;charset=utf-8"

	resp, err = http.Post(url, body_type, req)
	defer resp.Body.Close()
	if err != nil {
		logger.Info("sth wrong")
		logger.Info(err)
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
		logger.Info("cant unmarshal")
		logger.Error(err)
		return false
	}
	//logger.Info(res)
	if res.Errcode == 0 {
		logger.Info("nice")
		return true
	}
	logger.Info(res.Errmsg)
	return false
}
