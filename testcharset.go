package main

import (
	"fmt"
	"golang.org/x/text/transform"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
	"gopkg.in/iconv.v1"
	"github.com/op/go-logging"
	"github.com/PuerkitoBio/goquery"
)

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

var logger *logging.Logger

func init()  {
	logger = logging.MustGetLogger("main")
}

func main() {
	aConv, err := iconv.Open("utf-8", "GBK")
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer aConv.Close()

	resp, err := http.Get("http://www.aoshu.com/e/20180925/5baa00ca11932_2.shtml")
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	cont,_ := doc.Find("div.content").Find("p").Eq(1).GBKHtml()

	if err != nil {
		logger.Info(err)
	}

	defer resp.Body.Close()
	//input, err := ioutil.ReadAll(resp.Body)
	logger.Info(string(cont))
	//str := aConv.ConvString(string(input))
	//fmt.Println(str)
}