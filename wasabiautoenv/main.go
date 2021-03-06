package main

import (
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"
	"fmt"
	"myproj.lee/try/wasabiautoenv/models"
)

const (
	org = "baas1"
	ip = "192.168.9.83"
)

func main() {
	//simpleTest()

	//createUserMsp()
	//invoke()
}


func createIdentity()  {
	host1 := "http://192.168.9.82:8081/member/id"
	ciRequest := models.SetupRequest{
		Consensus:"etcdraft",
		PeerPorts:[]models.PeerPorts{
			models.PeerPorts{
				Main:30031,
				Chaincode:30032,
			},
		},
		Company:"baas2",
		AutoGeneratedCerts:true,
	}
	data, err := json.Marshal(ciRequest)
	if err != nil {
		panic(err)
	}

	wrt := bytes.NewBuffer(data)

	resp, err := http.Post(host1, "application/json", wrt)

	if err != nil {
		panic(err)
	}
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ret))
}

type CreateIndirectRequest struct {
	Orgname  string `json:"orgname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func createUserMsp()  {
	cir := CreateIndirectRequest{
		Orgname:"org1",
		Username:"user1",
		Password:"12345678",
	}

	data, err := json.Marshal(cir)
	if err != nil {
		panic(err)
	}

	wrt := bytes.NewBuffer(data)

	resp, err := http.Post("http://"+ip+":8081/indirect/create", "application/json", wrt)

	if err != nil {
		panic(err)
	}
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ret))
}



func invoke()  {

	iir := models.IndirectInvokeRequest{
		Userid:"b2c8c090-413e-4433-a5cc-941f62ec1ae7",
		Org:"org444",
		ChannelName:"channel1",
		CcName:"example1",
		Args:[]string{},
	}

	data, err := json.Marshal(iir)
	if err != nil {
		panic(err)
	}

	wrt := bytes.NewBuffer(data)

	resp, err := http.Post("http://"+ip+":8081/indirect/invoke", "application/json", wrt)

	if err != nil {
		panic(err)
	}
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ret))
}