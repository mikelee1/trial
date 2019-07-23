package main_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type Token struct {
	Token string `json:"token"`
}

func Test_Main(t *testing.T) {
	data := `{"username":"someone", "password":"p@ssword"}`
	resp, e := http.Post("http://127.0.0.1:8081/login", "application/json;", strings.NewReader(data))
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	tmp := &Token{}
	json.Unmarshal(body, tmp)
	fmt.Println(tmp.Token)
}
