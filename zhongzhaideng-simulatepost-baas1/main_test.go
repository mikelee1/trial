package main_test

import (
	"testing"
	"github.com/op/go-logging"
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"
	types2 "myproj.lee/try/wasabi-simulatepost-baas1/types"
	"fmt"
	"time"
	"os"
	"mime/multipart"
	"io"
	"path/filepath"
	"log"
)

var lresp = types2.LoginResp{}
var baas1 = "192.168.9.87"
var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("simulatepost")
}

var (
	version  = "1.0.18"
	ccName   = "ccdc"
	channel  = "mychannel"
	filename = "/Users/leemike/go/src/yindeng/yindeng.tar.gz"
)
//var filename = "/Users/leemike/go/src/wasabi/backEnd/vendor/github.com/hyperledger/fabric/examples/chaincode/go/marbles.tar.gz"
//------内网上传升级合约-----
func Test_install_upgrade_chaincode(t *testing.T) {
	time.Now()
	Login(baas1)
	Upload(baas1)
	UpgradeChaincode(baas1)
}

func Test_invoke_big_chaincode(t *testing.T) {
	time.Now()
	Login(baas1)
	InvokeBig(baas1)
}

//内网上传创建合约
func Test_install_instantiate_chaincode(t *testing.T) {
	time.Now()
	Login(baas1)
	Upload(baas1)
	InstantiateChaincode(baas1)
}

func Login(ip string) {
	l := types2.LoginReq{
		Username: "admin",
		Password: "yunphant",
	}
	lb, _ := json.Marshal(l)
	resp, err := http.Post(fmt.Sprintf("http://%s:8081/login", ip), "application/json", bytes.NewBuffer(lb))

	if err != nil {
		logger.Info(err)
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &lresp)
	if err != nil {
		logger.Info(err)
		panic(err)
	}
	logger.Info("token: ", lresp.Data.Token)
}

func Upload(ip string) {
	url := fmt.Sprintf("http://%s:8081/chaincode/upload", ip)

	filetype := "file"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(filetype, filepath.Base(file.Name()))

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(part, file)
	writer.Close()
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("authorization", lresp.Data.Token)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fr := FileResp{}
	err = json.Unmarshal(content, &fr)
	if err != nil {
		log.Fatal(err)
	}
	hash = fr.Hash
	fmt.Println(hash)
}

var hash = ""

func InstantiateChaincode(ip string) {

	l := fmt.Sprintf(`{"cc_hash":"%s",
	"cc_name":"%s","cc_version":"%s","metadata_path":"",
	"cc_path":"yindeng","args":["init"],"channel_name":"%s","cc_policy":"OR('baas1.member')","peer_nodes":["peer-0-baas1","peer-1-baas1"]}`, hash, ccName, version, channel)

	req, _ := http.NewRequest("POST", fmt.Sprintf("http://%s:8081/chaincode/installandinstantiate", ip), bytes.NewBuffer([]byte(l)))

	req.Header.Add("authorization", lresp.Data.Token)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		logger.Info(err)
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	logger.Info("res: ", string(body))
}

//带authorization的post请求
func UpgradeChaincode(ip string) {

	l := fmt.Sprintf(`{"cc_hash":"%s",
	"cc_name":"%s","cc_version":"%s",
	"cc_path":"yindeng","args":["init","1"],"channel_name":"%s","cc_policy":"OR('baas1.member')","peer_nodes":["peer-0-baas1","peer-1-baas1"]}`, hash, ccName, version, channel)

	req, _ := http.NewRequest("POST", fmt.Sprintf("http://%s:8081/chaincode/installandupgrade", ip), bytes.NewBuffer([]byte(l)))

	req.Header.Add("authorization", lresp.Data.Token)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		logger.Info(err)
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	logger.Info("res: ", string(body))
}

type FileResp struct {
	FileID   string
	FileName string
	Hash     string
}

type InvokeReq struct {
	Channel string
	Ccname  string
	Fcn     string
	Args    []string
}

func InvokeBig(ip string) {
	ir := InvokeReq{
		Channel: "channel1",
		Ccname:  ccName,
		Fcn:     "create_product",
	}
	file, err := ioutil.ReadFile("test.tar.gz")
	if err != nil {
		logger.Info(err)
		panic(err)
	}
	//fmt.Println(string(file))
	l := ReqProduct{
		Id:       "id19286",
		ProdName: "pppp6",
		Owner:    string(file),
	}
	lByte, _ := json.Marshal(l)
	ir.Args = append(ir.Args, string(lByte))
	irByte, _ := json.Marshal(ir)

	req, _ := http.NewRequest("POST", fmt.Sprintf("http://%s:8080/chaincode/invoke", ip), bytes.NewBuffer(irByte))
	req.Header.Add("authorization", lresp.Data.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Info(err)
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	logger.Info("res: ", string(body))
}

type ReqProduct struct {
	Id       string `json:"id"`
	ProdName string `json:"prod_name"`
	Owner    string `json:"owner"`
}
