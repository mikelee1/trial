package main_test

import (
	"testing"
	"github.com/op/go-logging"
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"
	types2 "myproj/try/wasabi-simulatepost-baas1/types"
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

//------内网上传升级合约-----
func Test_install_instantiate_chaincode(t *testing.T) {
	time.Now()
	Login(baas1)
	Upload(baas1)
	InstallChaincode(baas1)
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
	filename := "/Users/leemike/go/src/yindeng/yindeng.tar.gz"
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
//带authorization的post请求
func InstallChaincode(ip string) {

	l := fmt.Sprintf(`{"cc_hash":"%s",
	"cc_name":"yindeng21","cc_version":"1.11",
	"cc_path":"yindeng","args":["init"],"channel_name":"channel2","cc_policy":"OR('baas1.member')","peer_nodes":["peer-0-baas1"]}`, hash)

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
