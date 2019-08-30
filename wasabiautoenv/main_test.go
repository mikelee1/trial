package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/op/go-logging"
	"io/ioutil"
	"myproj/try/common/file"
	logger2 "myproj/try/common/logger"
	"myproj/try/wasabiautoenv/models"
	"net/http"
	"os"
	"testing"
	"github.com/appleboy/easyssh-proxy"
	"time"
)

var (
	logger         *logging.Logger
	inviteCodeFile = "./invitecode"
	baas1          = "baas1"
	baas2          = "baas4"
	baas1Host      = "192.168.9.83"
	baas2Host      = "192.168.9.87"

	channel    = "channel1"
	chaincode  = "example1"
	orgname    = "org2"
	username   = "user2"
	password   = "12345678"

	//ssh到baas1上的参数
	dirpath    = "/home/ubuntu"
	hostuser   = "ubuntu"

	indirectid = ""
	cchash     = "bdb2b28b8f83c06f594cd2ac20e2e126" //默认情况下examplecc的hash都一样
)

func init() {
	logger = logger2.GetLogger()
}
//baas1进行setup
func Test_Setup(t *testing.T) {
	org := baas1
	peers := []models.PeerPorts{
		models.PeerPorts{
			Main:      30031,
			Chaincode: 30032,
		},
		models.PeerPorts{
			Main:      30033,
			Chaincode: 30034,
		},
	}
	orderer0 := models.OrdererPorts{
		Main: 30020,
	}
	orderer1 := models.OrdererPorts{
		Main: 30022,
	}
	orderer2 := models.OrdererPorts{
		Main: 30024,
	}
	SetupRequest := models.SetupRequest{
		OrgName:      org,
		PeerPorts:    peers,
		OrdererPorts: []models.OrdererPorts{orderer0, orderer1, orderer2},
		Consensus:    "etcdraft",
	}

	data, err := json.Marshal(SetupRequest)
	if err != nil {
		panic(err)
	}

	wrt := bytes.NewBuffer(data)

	resp, err := http.Post("http://"+baas1Host+":8081/member/setup", "application/json", wrt)

	if err != nil {
		panic(err)
	}
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ret))
	//time.Sleep(50*time.Second)
}
//baas2创建identity -> baas1邀请、同意 -> baas2加入网络 -> 创建和加入通道
func Test_Main(t *testing.T) {
	Test_CreateIdentity(t)
	Test_ScpIdentity(t)
	Test_Invite(t)
	Test_AgreeInvitation(t)
	Test_InviteCode(t)
	Test_Join(t)
	Test_ChannelJoin(t)
	Test_CreateAndJoinChannel(t)
}

func Test_CreateIdentity(t *testing.T) {
	host1 := "http://" + baas2Host + ":8081/member/id"
	ciRequest := models.SetupRequest{
		Consensus: "etcdraft",
		PeerPorts: []models.PeerPorts{
			models.PeerPorts{
				Main:      30031,
				Chaincode: 30032,
			},
		},
		//OrdererPorts: []models.OrdererPorts{
		//	models.OrdererPorts{
		//		Main:  30021,
		//		Debug: 30022,
		//	},
		//},
		Company:            baas2,
		AutoGeneratedCerts: true,
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

func Test_ScpIdentity(t *testing.T) {
	sc := NewSshClient(baas2Host)
	stdout, _, _, err := sc.Run("cd "+dirpath+"/go/src/wasabi && ./getidentity.sh")
	if err != nil {
		logger.Error(err)
		return
	}
	//创建获取命令输出管
	logger.Info(string(stdout))

	sc = NewSshClient(baas1Host)
	stdout, _, _, err = sc.Run(fmt.Sprintf("cd %s/go/src/wasabi && "+
		"docker exec wasabi mkdir -p /wasabi/%s && "+
		"docker cp identity wasabi:/wasabi/%s", dirpath, baas2, baas2),
	)
	if err != nil {
		logger.Error(err)
		return
	}
	//创建获取命令输出管
	logger.Info(string(stdout))
}

func Test_Invite(t *testing.T) {
	resp, err := http.Get("http://" + baas1Host + ":8081/member/invitation/start?inviter=" + baas1 + "&invitee=" + baas2)

	if err != nil {
		panic(err)
	}
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ret))
}

func Test_AgreeInvitation(t *testing.T) {
	resp, err := http.Get("http://" + baas1Host + ":8081/member/invitation/vote?inviter=" + baas1 + "&invitee=" + baas2 + "&accept=true")
	if err != nil {
		t.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func Test_InviteCode(t *testing.T) {
	dir, _ := os.Getwd()
	logger.Info(dir)
	resp, err := http.Get("http://" + baas1Host + ":8081/member/ic")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	inviteCode, _ := ioutil.ReadAll(resp.Body)
	var f *os.File
	if file.CheckFileIsExist(inviteCodeFile) { //如果文件存在
		f, err = os.OpenFile(inviteCodeFile, os.O_TRUNC|os.O_WRONLY, 0666) //打开文件
	} else {
		f, err = os.Create(inviteCodeFile) //创建文件
	}

	if err != nil {
		logger.Error(err)
		return
	}
	defer f.Close()
	f.WriteString(string(inviteCode))
}

func Test_Join(t *testing.T) {
	data, err := ioutil.ReadFile(inviteCodeFile)
	if err != nil {
		t.Fatal(err)
	}
	wrt := bytes.NewBuffer(data)

	resp, err := http.Post("http://"+baas2Host+":8081/member/join", "application/json", wrt)

	if err != nil {
		t.Fatal(err)
	}
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(ret))
}

func Test_ChannelJoin(t *testing.T)  {
	cjRequest := models.JoinChannelRequest{
		Peers: []string{"peer-0-"+baas2},
		ChannelName:channel,
	}
	data, err := json.Marshal(cjRequest)
	if err != nil {
		t.Fatal(err)
	}
	wrt := bytes.NewBuffer(data)

	resp, err := http.Post("http://"+baas2Host+":8081/channel/join", "application/json", wrt)

	if err != nil {
		t.Fatal(err)
	}
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(ret))
}

func Test_CreateAndJoinChannel(t *testing.T) {
	data := models.InitChannelRequest{
		ChannelName: "channel1",
		Orgs:        []string{baas1, baas2},
		Peers:       []string{"peer-0-" + baas1, "peer-1-" + baas1},
	}
	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)

	_, err := http.Post("http://"+baas2Host+":8081/channel/createandjoin", "application/json", wrt)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_InstallAndInstantiate(t *testing.T) {
	data := models.InstallAndInstantiateChainCodeRequest{
		Args:        []string{"init", "a", "100", "b", "100"},
		CcHash:      cchash,
		CcName:      chaincode,
		CcPath:      "example_cc",
		CcVersion:   "1.0.0",
		ChannelName: channel,
		PeerNodes:   []string{"peer-0-"+baas2},
	}
	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)

	_, err := http.Post("http://"+baas2Host+":8081/chaincode/installandinstantiate", "application/json", wrt)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_InstallAndUpgrade(t *testing.T) {
	data := models.InstallAndInstantiateChainCodeRequest{
		Args:        []string{"init", "a", "100", "b", "100"},
		CcHash:      cchash,
		CcName:      chaincode,
		CcPath:      "example_cc",
		CcVersion:   "1.0.2",
		ChannelName: channel,
		PeerNodes:   []string{"peer-0-baas2"},
	}
	bytedata, _ := json.Marshal(data)
	wrt := bytes.NewBuffer(bytedata)
	_, err := http.Post("http://"+baas2Host+":8081/chaincode/installandupgrade", "application/json", wrt)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_CreateIndirect(t *testing.T) {
	data := models.CreateIndirectRequest{
		Orgname:  orgname,
		Username: username,
		Password: password,
	}
	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)

	_, err := http.Post("http://"+baas2Host+":8081/indirect/create", "application/json", wrt)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_AddAuth(t *testing.T) {
	indirectid = LoginIndirect()
	data := models.AuthUpdate{
		Indirectid:  indirectid,
		Items:map[string][]string{
			channel: []string{chaincode},
		},
	}
	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)

	_, err := http.Post("http://"+baas2Host+":8081/direct/console/auth/update", "application/json", wrt)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_LoginIndirect(t *testing.T) {
	id := LoginIndirect()
	logger.Info(id)
}

func LoginIndirect() string {
	data := models.LoginRequest{
		Username: username,
		Password: password,
	}
	bytedata, _ := json.Marshal(data)
	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://"+baas2Host+":8081/login", "application/json", wrt)
	if err != nil {
		return err.Error()
	}
	a, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	b := &models.LoginResp{}
	err = json.Unmarshal(a, b)
	if err != nil {
		logger.Error(err)
		return err.Error()
	}
	indirectid = b.UserId
	return indirectid
}

func Test_Invoke(t *testing.T) {
	indirectid = LoginIndirect()
	data := models.IndirectInvokeRequest{
		ChannelName: channel,
		CcName:      chaincode,
		Args:        []string{"move", "a", "b", "3"},
		Userid:      indirectid,
	}
	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)

	_, err := http.Post("http://"+baas2Host+":8081/indirect/invoke", "application/json", wrt)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Query(t *testing.T) {
	indirectid = LoginIndirect()
	data := models.IndirectInvokeRequest{
		ChannelName: channel,
		CcName:      chaincode,
		Args:        []string{"query", "a"},
		Userid:      indirectid,
	}
	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)

	resp, err := http.Post("http://"+baas2Host+":8081/indirect/query", "application/json", wrt)

	if err != nil {
		t.Fatal(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	t.Log(string(body))
}


func NewSshClient(server string) *easyssh.MakeConfig {
	ssh := &easyssh.MakeConfig{
		User:   hostuser,
		Server: server,
		KeyPath: "/Users/leemike/.ssh/id_rsa",//私钥
		Port:    "22",
		Timeout: 60 * time.Second,
	}

	_, _, _, err := ssh.Run("cd "+dirpath+"/go/src/wasabi && pwd && ls", 60*time.Second)
	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
		return nil
	} else {
		fmt.Println("connect ssh "+server+" ok")
	}
	return ssh
}

