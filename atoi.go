package main

import (
	"fmt"
	"github.com/hyperledger/fabric/sdk"
	"strconv"
	"time"
)

func main() {
	//s := "0101"
	//if strings.Split(s,"")[0]=="0"{
	//	fmt.Println("wrong")
	//	return
	//}
	//a,_ := strconv.Atoi(s)
	//fmt.Println(a)
	var b float64
	b = 19.2323
	c := strconv.FormatFloat(b, 'f', -1, 32)
	fmt.Println(c)
}

//service/chaincode/chaincode.go
func (cc *Chaincode) InvokeOrQuery(invoke bool, channelName string, peers []*sdk.Endpoint, orderers []*sdk.Endpoint, args [][]byte) (*pp.ProposalResponse, error) {
	client := cc.client
	ccName := cc.ccName

	txID, prop, resps, endorser, err := endorseOneOfList(client, channelName, ccName, args, nil, peers)
	if err != nil {
		logger.Error("Error endorsing", err)
		return nil, err
	}
	if invoke {
		err = broadcastOneOfList(client, prop, resps, orderers)
		if err != nil {
			logger.Error("Error broadcasing", err)
			return nil, err
		}

		valid, err := client.WaitTx(channelName, txID, endorser, WaitTxTimeout)
		if err != nil {
			logger.Error("Error waiting transaction", err)
			return resps[0], err
		}

		if !valid {
			return resps[0], errors.New("invoke is not valid, please try again")
		}
	}
	logger.Info("resps are:%v\n", resps)
	if len(resps) == 0 {
		return nil, errors.New("no proposal responses received - this might indicate a bug")
	}

	logger.Info("Successfully invoke  chaincode")
	return resps[0], err
}

//controllers/chaincodecontroller
func serviceNodesToEndpointList(serviceNodes []*chaincode.ServiceNode, timeout time.Duration, cert []byte) []*sdk.Endpoint {
	var endpoints []*sdk.Endpoint
	for _, sn := range serviceNodes {
		endpoints = append(endpoints, &sdk.Endpoint{
			Address:  sn.Endpoint,
			Override: "", // pay attention
			TLS:      cert,
			Timeout:  timeout,
		})
	}
	return endpoints
}
