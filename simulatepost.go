package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/sdk"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var ri int

func main() {
	rand.Seed(time.Now().UnixNano())

	ri = rand.Intn(1)
	fmt.Println(ri)
	var waitc chan bool
	res := SimulatePost("http://192.168.9.21:8089/chaincode/query")
	fmt.Println(res)
	<-waitc
}
func SimulatePost(url string) string {
	Tmpa := &sdk.ChaincodeInvokeSpec{}
	//Tmpa.Channel = "mychannel1"
	//Tmpa.CCName = "usercc"
	//Tmpa.Fcn="CreateOrder"
	//Tmpa.Args = []string{"9cbf2a34-4905-478b-9af8-decc0a4baded","5a93df22-b97a-4908-99cb-01933ee545a1","ccc"}
	//

	Tmpa.Channel = "mychannel1"
	Tmpa.CCName = "usercc"
	Tmpa.Fcn = "GetHistoryOrder"
	Tmpa.Args = []string{"eb4d0add0c72fef541247a04101f37d91092f8af", "9cbf2a34-4905-478b-9af8-decc0a4baded", "5a93df22-b97a-4908-99cb-01933ee545a1"}

	jsonValue, _ := json.Marshal(Tmpa)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
