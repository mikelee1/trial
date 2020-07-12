package main

import (
	"github.com/op/go-logging"
	logger2 "myproj.lee/try/common/logger"
	"net/http"
	"net/url"
	"time"
	"myproj.lee/try/fanqiang/models"
	"myproj.lee/try/common/fmtstruct"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

var logger *logging.Logger

const (
	host = "https://api.huobi.pro/market/history/kline?period=1min&size=2&symbol=btcusdt"
	wshost = "wss://api.huobi.pro/ws"
)

func init() {
	logger = logger2.GetLogger()
}

func main() {
	ReqConn(5)
}

func ReqConn(d int)  {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:1087")
	}

	transport := &http.Transport{Proxy: proxy}

	client := &http.Client{Transport: transport}
	tick := time.NewTicker(time.Duration(d)*time.Second)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			resp, err := client.Get(host)
			if err != nil {
				logger.Error(err)
				return
			}

			body, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()

			minfo := models.MarketInfo{}
			err = json.Unmarshal(body, &minfo)
			if err != nil {
				logger.Error(err)
				return
			}

			for _,v := range minfo.Data{
				logger.Info(v)
				i,_ := strconv.ParseInt(strconv.Itoa(int(v.Id)),10,64)
				v.Time = time.Unix(i,0).Format("01-02 15:04:05")
			}
			logger.Info(fmtstruct.String(minfo.Data))
		case <- time.After(time.Second*20):
			logger.Info("time out")
			break
		}
	}
	logger.Info("exit")

}

