package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/op/go-logging"
	"io/ioutil"
	"net/http"
	"strings"
	"vnt-candy-box/utils/phone"
)

var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("vnt")
}

func main() {
	//a,err := phone.ParseToInternational("5856533056","US")
	//fmt.Println(a,err)
	err := SendVerifyMessage("18724006865", "1231", "CN")
	logger.Debug(err)
}

func SendVerifyMessage(phoneNumber string, code string, country string) error {
	url := "http://" + "47.75.188.96" + ":" + "10001" + "/verifycode"
	var err error
	pn := phoneNumber
	// 国际号码支持
	if country != "CN" {
		pn, err = phone.ParseToInternational(phoneNumber, country)
		if err != nil {
			return err
		}
		logger.Debugf(pn)
		pn = "00" + strings.Replace(pn, "+", "", -1)
		logger.Debugf(pn)
	}
	payload, err := json.Marshal(map[string]string{
		"phoneNumber": pn,
		"code":        code,
		"country":     country,
	})
	if err != nil {
		logger.Errorf("Error encoding MessageSendingRequest : %s", err)
		return err
	}
	p := strings.NewReader(string(payload))
	logger.Debug(url, p)
	res, err := http.Post(url, "application/json", p)
	if err != nil {
		logger.Errorf("Error sending request to %s : %s", url, err)
		return err
	}
	defer res.Body.Close()
	logger.Debug(res)
	if res.StatusCode == 200 {
		return nil
	} else if res.StatusCode == 500 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			logger.Errorf("Error reading response body : %s", err)
			return err
		}
		mse := &MsgServiceError{}
		// 可能出现短信服务超时的情况，返回500，但是response为空，且短信发出
		if len(body) == 0 {
			logger.Warningf("No response from message service . Check the message service log for timeout error.")
			return nil
		}
		if err := json.Unmarshal(body, mse); err != nil {
			logger.Errorf("Error unmarshalling MsgServiceError : %s , body: %s", err, string(body))
			return err
		}
		err = errors.New(fmt.Sprintf("Error sending verify message [%s] to [%s] : %v", code, phoneNumber, mse))
		logger.Error(err.Error())
		return err
	} else {
		err := errors.New(fmt.Sprintf("Unhandled error response [%d] : %s", res.StatusCode, res))
		logger.Error(err.Error())
		return err
	}
}

type MsgServiceError struct {
	Message   string `json:"Message"`
	RequestID string `json:"RequestId"`
	Code      string `json:"Code"`
}
