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

//------内网setup\生成identity-----
func Test_init(t *testing.T) {
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

//func postFile(url, filename, filePath string) error {
//
//	//打开文件句柄操作
//	file, err := os.Open(filePath)
//	if err != nil {
//		fmt.Println("error opening file")
//		return err
//	}
//	defer file.Close()
//
//	//创建一个模拟的form中的一个选项,这个form项现在是空的
//	bodyBuf := &bytes.Buffer{}
//	bodyWriter := multipart.NewWriter(bodyBuf)
//
//	//关键的一步操作, 设置文件的上传参数叫uploadfile, 文件名是filename,
//	//相当于现在还没选择文件, form项里选择文件的选项
//	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
//	if err != nil {
//		fmt.Println("error writing to buffer")
//		return err
//	}
//
//	//iocopy 这里相当于选择了文件,将文件放到form中
//	_, err = io.Copy(fileWriter, file)
//	if err != nil {
//		panic(err)
//		return err
//	}
//
//	//获取上传文件的类型,multipart/form-data; boundary=...
//	contentType := bodyWriter.FormDataContentType()
//	//fmt.Println(contentType)
//
//	//这个很关键,必须这样写关闭,不能使用defer关闭,不然会导致错误
//	bodyWriter.Close()
//
//	//fmt.Println("lresp.Data.Token: ", lresp.Data.Token)
//	err = bodyWriter.WriteField("Authorization", lresp.Data.Token)
//	if err != nil {
//		panic(err)
//	}
//	//这里就是上传的其他参数设置,可以使用 bodyWriter.WriteField(key, val) 方法
//	//也可以自己在重新使用  multipart.NewWriter 重新建立一项,这个再server 会有例子
//	params := map[string]string{
//		"filename": filename,
//	}
//	//这种设置值得仿佛 和下面再从新创建一个的一样
//	for key, val := range params {
//		_ = bodyWriter.WriteField(key, val)
//	}
//	fmt.Println(bodyBuf)
//	h := http.Header{}
//
//	//发送post请求到服务端
//	resp, err := http.Post(url, contentType, bodyBuf)
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//	defer resp.Body.Close()
//	resp_body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//	fmt.Println(resp.Status)
//	fmt.Println(string(resp_body))
//	return nil
//}

func postFile(url string, filename string, filetype string) []byte {
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
	fmt.Println(string(content))
	fr := FileResp{}
	err = json.Unmarshal(content, &fr)
	if err != nil {
		log.Fatal(err)
	}
	hash = fr.Hash
	return content
}

type FileResp struct {
	FileID   string
	FileName string
	Hash     string
}
