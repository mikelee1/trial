package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/sdk"
	"os"
	"path"
)

const (
	org    = "org1"
	mspDir = "./msp"
	baas   = "baas4"
	user   = "user1"
	//chainid = "channel1"
	//chaincode = "example1"
)

var err error

func getCA(dir string, msp string) (*sdk.CA, error) {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return sdk.NewCA(dir, msp)
	}
	if err == nil {
		if !info.IsDir() {
			return nil, errors.New("msp path is not a directory, but a file")
		}
		return sdk.ConstructCAFromDir(dir)
	}
	return nil, err
}

func CreateCa(Causer string) (bool, error, string,*sdk.CA) {
	orgCA, err := getCA(path.Join(mspDir, baas), baas)
	b := []string{Causer}
	if err = orgCA.GenerateMSP(nil, b); err != nil {
		fmt.Errorf("Error generating msp:%s ", err.Error())
	}
	return true, nil, path.Join(mspDir, baas, "users", Causer, "msp"),orgCA
}

func simpleTest() {
	var err error
	_, err, capath,orgCA := CreateCa(user)
	if err != nil {
		fmt.Errorf("error: %s", err.Error())
	}


	//cli, err := sdk.NewClient(orgCA.AdminCommonName(), baas, orgCA.AdminMSPDir(), false)
	//if err != nil {
	//	fmt.Errorf("error: %s", err.Error())
	//}

	cli, err := sdk.NewClient(org, baas, capath, true)
	if err != nil {
		fmt.Errorf("error: %s", err.Error())
	}

	cli, err = sdk.NewClient(orgCA.AdminCommonName(), baas, orgCA.AdminMSPDir(), true)
	if err != nil {
		fmt.Errorf("error: %s", err.Error())
	}

	fmt.Println(cli, err)

}

func main() {
	simpleTest()
}
