package main_test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	models "myproj/try/testwasabi/model"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type NewCreateChannelRequest struct {
	Orgs        []string
	ChannelName string
}

//创建channel,注意no baas2
func TestCreateChannel(t *testing.T) {
	channelname := "testchannel1"
	orgs := []string{"baas1"}
	ccr := &NewCreateChannelRequest{
		Orgs:        orgs,
		ChannelName: channelname,
	}

	data, err := json.Marshal(ccr)
	if err != nil {
		t.Fatal(err)
	}
	wrt := bytes.NewBuffer(data)

	resp, err := http.Post("http://192.168.9.21:8080/channel/create", "application/json", wrt)
	if err != nil {
		t.Fatal(err)
	}
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(ret))
}

type JoinChannelRequest struct {
	Peers       []string
	ChannelName string
}

//加入链
func TestJoinChannel(t *testing.T) {
	channelname := "testchannel1"
	peers := []string{"peer-0-baas1"}

	jcr := &JoinChannelRequest{
		Peers:       peers,
		ChannelName: channelname,
	}

	data, err := json.Marshal(jcr)
	if err != nil {
		t.Fatal(err)
	}
	wrt := bytes.NewBuffer(data)

	resp, err := http.Post("http://192.168.9.21:8080/channel/join", "application/json", wrt)
	if err != nil {
		t.Fatal(err)
	}
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(ret))
}

func TestIdentity(t *testing.T) {
	peers := []*models.ServiceNode{
		&models.ServiceNode{
			ID:               "peer0_lily",
			Endpoint:         "192.168.9.21:31010",
			ExternalEndpoint: "192.168.9.21:31010",
			Public:           true,
		},
	}

	orderers := []*models.ServiceNode{
		&models.ServiceNode{
			ID:               "orderer0_lily",
			Endpoint:         "192.168.9.21:31020",
			ExternalEndpoint: "192.168.9.21:31020",
			Public:           true,
		},
		&models.ServiceNode{
			ID:               "orderer1_lily",
			Endpoint:         "192.168.9.21:31030",
			ExternalEndpoint: "192.168.9.21:31030",
			Public:           true,
		},
	}
	orgs := []*models.OrgInfo{
		&models.OrgInfo{
			OrgName:      "baas3",
			OrgMSP:       "baas3",
			MspID:        "baas3",
			PeerNodes:    peers,
			OrdererNodes: orderers,
		},
	}

	idreq := &models.IdentityRequest{
		Orgs: orgs,
	}

	data, err := json.Marshal(idreq)
	if err != nil {
		t.Fatal(err)
	}

	wrt := bytes.NewBuffer(data)

	resp, err := http.Post("http://192.168.9.21:8080/channel/invite/identity", "application/json", wrt)

	if err != nil {
		t.Fatal(err)
	}
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile("orgIdentity", ret, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(t *testing.T) {
	url := "http://192.168.9.21:8080/login"
	lr, err := json.Marshal(&LoginRequest{"admin", "yunphant"})
	if err != nil {
		t.Fatal(err.Error())
	}
	lrt := bytes.NewBuffer(lr)
	resp, err := http.Post(url, "application/json", lrt)
	if err != nil {
		t.Fatal(err.Error())
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("login fail!", err)
	}
	t.Log(string(result))
}
func TestUpload(t *testing.T) {
	login(t)
	time.Sleep(time.Second * 3)
	url := "http://192.168.9.21:8080/chaincode/upload"
	filePath := "./chaincode/example.tar.gz"
	req, err := newUploadRequest(url, nil, "file", filePath)
	if err != nil {
		t.Fatalf(err.Error())
		return
	}

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	t.Log(res)
	t.Log(string(body))
}

// 新建上传请求
func newUploadRequest(link string, params map[string]string, name, path string) (*http.Request, error) {
	fp, err := os.Open(path) // 打开文件句柄
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	body := &bytes.Buffer{}             // 初始化body参数
	writer := multipart.NewWriter(body) // 实例化multipart

	//<<<<<<< HEAD
	//	resp, err := http.Post("http://192.168.9.21:8080/chaincode/install", "application/json", wrt)
	//=======
	part, err := writer.CreateFormFile(name, filepath.Base(path)) // 创建multipart 文件字段
	//>>>>>>> ffb6345bdae9f4543216a1cf0d067f636dbb4d43
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, fp) // 写入文件数据到multipart
	for key, val := range params {
		_ = writer.WriteField(key, val) // 写入body中额外参数，比如七牛上传时需要提供token
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", link, body) // 新建请求
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "multipart/form-data; boundary="+writer.Boundary()) // 设置请求头,!!!非常重要，否则远端无法识别请求
	return req, nil
}

type InstallAndInstantiateChainCodeRequest struct {
	CcName      string   `json:"cc_name"`
	CcVersion   string   `json:"cc_version"`
	Args        []string `json:"args"`
	ChannelName string   `json:"channel_name"`
	PeerNodes   []string `json:"peer_nodes"`
	CcPath      string   `json:"cc_path"`
	CcHash      string   `json:"cc_hash"`
	Org         string
}

func TestInstallAndInstantiateChaincode(t *testing.T) {

	iitcq := &InstallAndInstantiateChainCodeRequest{
		CcName:      "example",
		CcVersion:   "1.0",
		Args:        []string{"init", "a", "1000000", "b", "1000000"},
		ChannelName: "testchannel1",
		PeerNodes:   []string{"peer-0-baas1"},
		CcPath:      "example_cc",
		CcHash:      "bdb2b28b8f83c06f594cd2ac20e2e126",
	}

	data, err := json.Marshal(iitcq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
	wrt := bytes.NewBuffer(data)

	resp, err := http.Post("http://192.168.9.21:8080/chaincode/installandinstantiate", "application/json", wrt)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(ret))
}

func TestMoveInvoke(t *testing.T) {
	org := "baas1"
	channelName := "testchannel1"
	ccName := "example"
	args := [][]byte{
		[]byte("move"),
		[]byte("a"),
		[]byte("b"),
		[]byte("1"),
	}
	peernodes := []*models.ServiceNode{
		&models.ServiceNode{
			ID:               "peer0_lily",
			Endpoint:         "192.168.9.21:30010",
			ExternalEndpoint: "192.168.9.21:30010",
			Public:           true,
		},
	}

	ordernodes := []*models.ServiceNode{
		&models.ServiceNode{
			ID:               "orderer0_lily",
			Endpoint:         "192.168.9.21:30020",
			ExternalEndpoint: "192.168.9.21:30020",
			Public:           true,
		},
		&models.ServiceNode{
			ID:               "orderer1_lily",
			Endpoint:         "192.168.9.21:30030",
			ExternalEndpoint: "192.168.9.21:30030",
			Public:           true,
		},
	}

	icr := &models.InvokeRequest{
		Invoke:      true,
		Org:         org,
		CcName:      ccName,
		ChannelName: channelName,
		Args:        args,

		PeerNodes:    peernodes,
		OrdererNodes: ordernodes,
	}

	data, err := json.Marshal(icr)
	if err != nil {
		t.Fatal(err)
	}
	wrt := bytes.NewBuffer(data)

	resp, err := http.Post("http://192.168.9.21:8080/chaincode/invoke", "application/json", wrt)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(ret))
}

func TestQueryInvoke(t *testing.T) {
	org := "baas1"
	channelName := "testchannel1"
	ccName := "example"
	args := [][]byte{
		[]byte("query"),
		[]byte("a"),
	}
	peernodes := []*models.ServiceNode{
		&models.ServiceNode{
			ID:               "peer0_lily",
			Endpoint:         "192.168.9.21:30010",
			ExternalEndpoint: "192.168.9.21:30010",
			Public:           true,
		},
	}

	ordernodes := []*models.ServiceNode{
		&models.ServiceNode{
			ID:               "orderer0_lily",
			Endpoint:         "192.168.9.21:30020",
			ExternalEndpoint: "192.168.9.21:30020",
			Public:           true,
		},
		&models.ServiceNode{
			ID:               "orderer1_lily",
			Endpoint:         "192.168.9.21:30030",
			ExternalEndpoint: "192.168.9.21:30030",
			Public:           true,
		},
	}

	icr := &models.InvokeRequest{
		Invoke:      false,
		Org:         org,
		CcName:      ccName,
		ChannelName: channelName,
		Args:        args,

		PeerNodes:    peernodes,
		OrdererNodes: ordernodes,
	}

	data, err := json.Marshal(icr)
	if err != nil {
		t.Fatal(err)
	}
	wrt := bytes.NewBuffer(data)

	resp, err := http.Post("http://192.168.9.21:8080/chaincode/invoke", "application/json", wrt)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(ret))
}

func TestQuery(t *testing.T) {
	org := "baas1"
	channelName := "testchannel1"
	ccName := "mycc3"
	args := []string{
		"query", "a",
	}
	icr := &models.IndirectInvokeRequest{
		Org:         org,
		CcName:      ccName,
		ChannelName: channelName,
		Args:        args,
	}

	data, err := json.Marshal(icr)
	if err != nil {
		t.Fatal(err)
	}
	wrt := bytes.NewBuffer(data)

	resp, err := http.Post("http://192.168.9.21:8080/indirect/query", "application/json", wrt)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(ret))
}

func TestQueryRaw(t *testing.T) {
	org := "baas1"
	channelName := "testchannel1"
	ccName := "mycc3"
	args := [][]byte{
		[]byte("query"),
		[]byte("a"),
	}
	peernodes := []*models.ServiceNode{
		&models.ServiceNode{
			ID:               "peer0_lily",
			Endpoint:         "192.168.9.21:30010",
			ExternalEndpoint: "192.168.9.21:30010",
			Public:           true,
		},
	}

	ordernodes := []*models.ServiceNode{
		&models.ServiceNode{
			ID:               "orderer0_lily",
			Endpoint:         "192.168.9.21:30020",
			ExternalEndpoint: "192.168.9.21:30020",
			Public:           true,
		},
		&models.ServiceNode{
			ID:               "orderer1_lily",
			Endpoint:         "192.168.9.21:30030",
			ExternalEndpoint: "192.168.9.21:30030",
			Public:           true,
		},
	}

	icr := &models.InvokeRequest{
		Invoke:      false,
		Org:         org,
		CcName:      ccName,
		ChannelName: channelName,
		Args:        args,

		PeerNodes:    peernodes,
		OrdererNodes: ordernodes,
	}

	data, err := json.Marshal(icr)
	if err != nil {
		t.Fatal(err)
	}
	wrt := bytes.NewBuffer(data)

	resp, err := http.Post("http://192.168.9.21:8080/indirect/queryraw", "application/json", wrt)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(ret))
}
