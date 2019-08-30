package main_test

import (
	"bytes"
	"encoding/json"
	"github.com/op/go-logging"
	"goSrcRead/fmt"
	"io/ioutil"
	"jiaoan/common/constant"
	"jiaoan/services/account_center/protocol/account"
	"jiaoan/services/admin_center/protocol/admin"
	"jiaoan/services/jiaoan_center/protocol/jiaoan"
	"net/http"
	"sync"
	"testing"
	"jiaoan/services/user_center/protocol/user"
)

var logger = logging.MustGetLogger("testwebserver")
var wg *sync.WaitGroup
var phone = "mike"
var password = "pass"
var newpassword = "123456"
var vericode string

func Test_all(t *testing.T) {
	Test_registe(t)
	Test_Login(t)
	Test_reset(t)
}

func Test_registe(t *testing.T) {
	var err1 error
	vericode, err1 = getverifycode(false, constant.ForRegiste, phone)
	if err1 != nil {
		logger.Error(err1.Error())
		return
	}
	registe(false, phone, password)
}

func Test_Login(t *testing.T) {
	vericode, _ = getverifycode(false, constant.ForLogin, phone)
	login(false, phone, password, vericode)
}

func Test_AdminLogin(t *testing.T) {
	vericode, _ = getverifycode(false, constant.ForLogin, phone)
	adminlogin(false, phone, password, vericode)
}

func Test_MultiLogin(t *testing.T) {
	wg = &sync.WaitGroup{}
	for _ = range [10]int{} {
		wg.Add(1)
		go login(true, phone, password, vericode)
	}
	wg.Wait()
}

func Test_checkverifycode(t *testing.T) {
	vericode, _ = getverifycode(false, constant.ForLogin, phone)
	checkverifycode(phone, constant.ForLogin, vericode)
}

func Test_reset(t *testing.T) {
	token := login(false, phone, password, "")
	resetpassword(token, password, newpassword)
	resetpassword(token, newpassword, password)
}

func Test_Getten(t *testing.T) {
	gettenjiaoan()
}

func Test_AdminEcho(t *testing.T) {
	echo(false)
}

func Test_DeleteJiaoan(t *testing.T) {
	deletejiaoan()
}

func Test_GetOneJiaoan(t *testing.T) {
	getonejiaoan()
}

func Test_Gettenprincipal(t *testing.T) {
	gettenprincipal()
}

func Test_Createprincipaljiaoan(t *testing.T) {
	Createprincipaljiaoan()
}

func Test_Deleteprincipaljiaoan(t *testing.T) {
	Deleteprincipaljiaoan()
}

func Test_RelateTeacher(t *testing.T)  {
	RelateTeacher()
}

func Test_DeRelateTeacher(t *testing.T)  {
	DeRelateTeacher()
}

func adminlogin(withWait bool, phone, password, vericode string) string {
	if withWait {
		defer wg.Done()
	}
	data := adminCenterProto.AdminLoginRequest{
		Phone:      phone,
		Password:   password,
		Verifycode: vericode,
	}
	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://127.0.0.1:8080/admin/login", "application/json",
		wrt)
	if err != nil {
		logger.Info(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info(err)
		return ""
	}
	logger.Info(string(body))
	res := &adminCenterProto.AdminLoginResponse{}
	json.Unmarshal(body, res)
	return res.Token
}


func login(withWait bool, phone, password, vericode string) string {
	if withWait {
		defer wg.Done()
	}
	data := accountCenterProto.LoginRequest{
		Phone:      phone,
		Password:   password,
		Verifycode: vericode,
	}
	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://127.0.0.1:8080/account/login", "application/json",
		wrt)
	if err != nil {
		logger.Info(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info(err)
		return ""
	}
	logger.Info(string(body))
	res := &accountCenterProto.LoginResponse{}
	json.Unmarshal(body, res)
	return res.Token
}

func echo(withWait bool) {
	if withWait {
		defer wg.Done()
	}
	data := adminCenterProto.EchoRequest{
		Ping: "mike",
	}
	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://127.0.0.1:8080/admin/echo", "application/json",
		wrt)
	if err != nil {
		logger.Info(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info(err)
		return
	}
	logger.Info(string(body))
}

//提供手机号和用途，获取校验码
func getverifycode(withWait bool, foruse constant.ForUseType, phone string) (string, error) {
	if withWait {
		defer wg.Done()
	}
	data := accountCenterProto.GetVerifyCodeRequest{
		Phone:  phone,
		Foruse: int32(foruse),
	}
	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://127.0.0.1:8080/account/verifycode/get", "application/json",
		wrt)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info(err)
		return "", err
	}
	res := &accountCenterProto.GetVerifyCodeResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		logger.Info(err)
		return "", err
	}
	if res.ResultMsg != "" {
		return "", fmt.Errorf(res.ResultMsg)
	}
	return res.Verifycode, nil
}

//提供手机号和校验码，验证
func checkverifycode(phone string, foruse constant.ForUseType, vericode string) {

	data := accountCenterProto.CheckVerifyCodeRequest{
		Phone:      phone,
		Foruse:     int32(foruse),
		Verifycode: vericode,
	}
	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://127.0.0.1:8080/account/verifycode/check", "application/json",
		wrt)
	if err != nil {
		logger.Info(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info(err)
		return
	}
	logger.Info(string(body))
}

func registe(withWait bool, phone, password string) {
	if withWait {
		defer wg.Done()
	}
	data := accountCenterProto.RegisteRequest{
		Phone:      phone,
		Password:   password,
		Verifycode: vericode,
	}
	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://127.0.0.1:8080/account/registe", "application/json",
		wrt)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info(err)
		return
	}
	res := &accountCenterProto.RegisteResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		logger.Info(err)
		return
	}
	logger.Info(res.Accountid)
}

func resetpassword(token, oldpassword, newpassword string) {
	data := &accountCenterProto.ResetPasswordRequest{
		Token:       token,
		Oldpassword: oldpassword,
		Newpassword: newpassword,
	}
	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://127.0.0.1:8080/account/password/reset", "application/json",
		wrt)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info(err)
		return
	}
	res := &accountCenterProto.ResetPasswordResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		logger.Info(err)
		return
	}
	logger.Info(res)
}

func gettenjiaoan() {
	data := &jiaoanCenterProto.GettenRequest{
		Pagenum:  0,
		Pagesize: 10,
	}

	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://127.0.0.1:8080/jiaoan/getten", "application/json",
		wrt)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info(err)
		return
	}
	res := &jiaoanCenterProto.GettenResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		logger.Info(err)
		return
	}
	logger.Info(res)
}

func deletejiaoan() {
	data := &jiaoanCenterProto.DeleteJiaoanRequest{
		Jiaoanid: 15,
	}

	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://127.0.0.1:8080/jiaoan/delete", "application/json",
		wrt)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info(err)
		return
	}
	res := &jiaoanCenterProto.DeleteJiaoanResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		logger.Info(err)
		return
	}
	logger.Info(res)
}

func getonejiaoan() {
	data := &jiaoanCenterProto.GetoneRequest{
		Jiaoanid: 16,
	}

	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://127.0.0.1:8080/jiaoan/getone", "application/json",
		wrt)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info(err)
		return
	}
	res := &jiaoanCenterProto.GetoneResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		logger.Info(err)
		return
	}
	logger.Info(res)
}


func gettenprincipal() {
	data := &adminCenterProto.GettenPrincipalRequest{
		Pagenum:  0,
		Pagesize: 10,
	}

	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://127.0.0.1:8080/admin/gettenprincipal", "application/json",
		wrt)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info(err)
		return
	}
	res := &adminCenterProto.GettenPrincipalResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		logger.Info(err)
		return
	}
	logger.Info(res)
}

func Createprincipaljiaoan() {
	data := &adminCenterProto.CreatePrincipalJiaoanRequest{
		Principalid:  1,
		Jiaoanid: 19,
	}

	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://127.0.0.1:8080/admin/createprincipaljiaoan", "application/json",
		wrt)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info(err)
		return
	}
	res := &adminCenterProto.CreatePrincipalJiaoanResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		logger.Info(err)
		return
	}
	logger.Info(res)
}

func Deleteprincipaljiaoan() {
	data := &adminCenterProto.DeletePrincipalJiaoanRequest{
		Principaljiaoanid:2,
		Token:"",
	}

	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://127.0.0.1:8080/admin/deleteprincipaljiaoan", "application/json",
		wrt)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info(err)
		return
	}
	res := &adminCenterProto.DeletePrincipalJiaoanResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		logger.Info(err)
		return
	}
	logger.Info(res)
}

func RelateTeacher() {
	data := &userCenterProto.RelateTeacherRequest{
		Principalid:1,
		Teacherid:2,
		Token:"",
	}

	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://127.0.0.1:8080/user/relate/teacher", "application/json",
		wrt)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info(err)
		return
	}
	res := &userCenterProto.RelateTeacherResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		logger.Info(err)
		return
	}
	logger.Info(res)
}

func DeRelateTeacher() {
	data := &userCenterProto.DeRelateTeacherRequest{
		Principalid:1,
		Teacherid:2,
		Token:"",
	}

	bytedata, _ := json.Marshal(data)

	wrt := bytes.NewBuffer(bytedata)
	resp, err := http.Post("http://127.0.0.1:8080/user/derelate/teacher", "application/json",
		wrt)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Info(err)
		return
	}
	res := &userCenterProto.DeRelateTeacherResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		logger.Info(err)
		return
	}
	logger.Info(res)
}