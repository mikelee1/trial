package main

import (
	"bytes"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/op/go-logging"
)
var logger = logging.MustGetLogger("testwasabi")
var baseurl = "http://192.168.9.21:8080"
var contenttype = "application/json;charset=utf-8"

var globaldirectsession string
var globalindirectsession string
func main() {
	DirectLogin()
	InDirectLogin()
	//TestOrg()
	//UserjoinApply()
	//UserjoinGetlist()
	//Monitorall()

}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func DirectLogin() {
	u := LoginRequest{Username:"admin",Password:"222"}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)

	resp, _ := http.Post(baseurl+"/login",contenttype,b)

	defer resp.Body.Close()
	body,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	globaldirectsession = string(body)
}

func InDirectLogin() {
	u := LoginRequest{Username:"333",Password:"333"}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)

	resp, _ := http.Post(baseurl+"/login",contenttype,b)
	defer resp.Body.Close()
	body,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	globalindirectsession = string(body)
}
