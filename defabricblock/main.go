package main

import (
	"github.com/op/go-logging"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/peer"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"fmt"
)

var logger1 *logging.Logger

var blockstring = "\n\xc1\a\ng\b\x03\x1a\f\b\x84\x93\xe2\xef\x05\x10\xc3\xe2\xc0\xd0\x03\"\achannel*@7c037007e867563dedf22ecf487781481e3de46439c3bb94a6a661a34bb61951:\n\x12\b\x12\x06sdzyb6\x12\xd5\x06\n\xb8\x06\n\x05baas1\x12\xae\x06-----BEGIN CERTIFICATE-----\nMIICLTCCAdOgAwIBAgIQD0X5dTQ2idoxmxCiYHim0zAKBggqhkjOPQQDAjB2MQsw\nCQYDVQQGEwJDTjERMA8GA1UECBMIWmhlamlhbmcxETAPBgNVBAcTCEhhbmd6aG91\nMRAwDgYDVQQJEwdjb21wYW55MQ8wDQYDVQQREwYzMTAwMDAxDjAMBgNVBAoTBWJh\nYXMxMQ4wDAYDVQQDEwViYWFzMTAeFw0xOTEyMTQwMjA2MDBaFw0yOTEyMTEwMjA2\nMDBaMGwxCzAJBgNVBAYTAkNOMREwDwYDVQQIEwhaaGVqaWFuZzERMA8GA1UEBxMI\nSGFuZ3pob3UxEDAOBgNVBAkTB2NvbXBhbnkxDzANBgNVBBETBjMxMDAwMDEUMBIG\nA1UEAwwLQWRtaW5AYmFhczEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS6lsmm\n084ahN541MhroSZh9PmhZOkjzdKqz52AMF+chVLvjpz6KjJqbTO+GwR9+MjFdr20\njt527mlCxX/li5rqo00wSzAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/BAIwADAr\nBgNVHSMEJDAigCBc5Qe/YfgxxVuZqsmP9WE5GYNq6iYzgmR1KQkeLJpwOTAKBggq\nhkjOPQQDAgNIADBFAiEAyhrcZUdU2cVtKbsB0ijyXMabNic6/tF3KzQlkFBPA6EC\nIDji15E0kpzYx8fe18P2f+DGcY0f/B183GrTPGc1Slvc\n-----END CERTIFICATE-----\n\x12\x18\x1b\x93\a\r\x17\xf1x\x89ļ\x83\x13\xbe\xdf\xf7\xdc9\xfc\xee-\xad)g\xb0\x12\xb1\x12\n\xae\x12\n\xd5\x06\n\xb8\x06\n\x05baas1\x12\xae\x06-----BEGIN CERTIFICATE-----\nMIICLTCCAdOgAwIBAgIQD0X5dTQ2idoxmxCiYHim0zAKBggqhkjOPQQDAjB2MQsw\nCQYDVQQGEwJDTjERMA8GA1UECBMIWmhlamlhbmcxETAPBgNVBAcTCEhhbmd6aG91\nMRAwDgYDVQQJEwdjb21wYW55MQ8wDQYDVQQREwYzMTAwMDAxDjAMBgNVBAoTBWJh\nYXMxMQ4wDAYDVQQDEwViYWFzMTAeFw0xOTEyMTQwMjA2MDBaFw0yOTEyMTEwMjA2\nMDBaMGwxCzAJBgNVBAYTAkNOMREwDwYDVQQIEwhaaGVqaWFuZzERMA8GA1UEBxMI\nSGFuZ3pob3UxEDAOBgNVBAkTB2NvbXBhbnkxDzANBgNVBBETBjMxMDAwMDEUMBIG\nA1UEAwwLQWRtaW5AYmFhczEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS6lsmm\n084ahN541MhroSZh9PmhZOkjzdKqz52AMF+chVLvjpz6KjJqbTO+GwR9+MjFdr20\njt527mlCxX/li5rqo00wSzAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/BAIwADAr\nBgNVHSMEJDAigCBc5Qe/YfgxxVuZqsmP9WE5GYNq6iYzgmR1KQkeLJpwOTAKBggq\nhkjOPQQDAgNIADBFAiEAyhrcZUdU2cVtKbsB0ijyXMabNic6/tF3KzQlkFBPA6EC\nIDji15E0kpzYx8fe18P2f+DGcY0f/B183GrTPGc1Slvc\n-----END CERTIFICATE-----\n\x12\x18\x1b\x93\a\r\x17\xf1x\x89ļ\x83\x13\xbe\xdf\xf7\xdc9\xfc\xee-\xad)g\xb0\x12\xd3\v\n\x9a\x01\n\x97\x01\n\x94\x01\b\x01\x12\b\x12\x06sdzyb6\x1a\x85\x01\n\rcreate_member\nt{\"ID\":\"1000002\",\"Name\":\"用户1\",\"TypeName\":1,\"CreateTime\":\"2006-01-02 15:04:06\",\"UpdateTime\":\"2006-01-02 15:04:07\"}\x12\xb3\n\n\xa4\x03\n E\x16\xc8\x1e\x11\x8aJ\x12R/{\xbeܑ\xa1\xb8\x80\xdc\xf0\xc1\xd51c\xb4\"s\x00\xdeϲ]\xef\x12\xff\x02\n\xf7\x01\x12\x16\n\x04lscc\x12\x0e\n\f\n\x06sdzyb6\x12\x02\b\"\x12\xdc\x01\n\x06sdzyb6\x12\xd1\x01\n\t\n\a1000002\x1a\xc3\x01\n\x15\x00zyb-member-\x001000002\x00\x1a\xa9\x01{\"ID\":\"1000002\",\"Name\":\"用户1\",\"TypeID\":\"\",\"TypeName\":1,\"CreateTime\":\"2006-01-02T15:04:06Z\",\"UpdateTime\":\"2006-01-02T15:04:07Z\",\"DelayAccountMoney\":0,\"AccountMoney\":0}\x12m\n\x06sdzyb6\x12@7c037007e867563dedf22ecf487781481e3de46439c3bb94a6a661a34bb61951\x1a\x03666\"\x1cthis is createmember payload\x1a\x03\b\xc8\x01\"\x0f\x12\x06sdzyb6\x1a\x051.0.0\x12\x89\a\n\xbc\x06\n\x05baas1\x12\xb2\x06-----BEGIN CERTIFICATE-----\nMIICLzCCAdWgAwIBAgIRAPkKdbfEyaqU8wNqYg2ggJQwCgYIKoZIzj0EAwIwdjEL\nMAkGA1UEBhMCQ04xETAPBgNVBAgTCFpoZWppYW5nMREwDwYDVQQHEwhIYW5nemhv\ndTEQMA4GA1UECRMHY29tcGFueTEPMA0GA1UEERMGMzEwMDAwMQ4wDAYDVQQKEwVi\nYWFzMTEOMAwGA1UEAxMFYmFhczEwHhcNMTkxMjE0MDIwNjAwWhcNMjkxMjExMDIw\nNjAwWjBtMQswCQYDVQQGEwJDTjERMA8GA1UECBMIWmhlamlhbmcxETAPBgNVBAcT\nCEhhbmd6aG91MRAwDgYDVQQJEwdjb21wYW55MQ8wDQYDVQQREwYzMTAwMDAxFTAT\nBgNVBAMTDHBlZXItMC1iYWFzMTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABM0/\nOnheZEilU3hWUNipZHc0KPEjQz30N+UsKw6FSak4gmXNQ6oecWrvv3YoMIYnZINx\n3gobSdpUJFTZ9zNY1pajTTBLMA4GA1UdDwEB/wQEAwIHgDAMBgNVHRMBAf8EAjAA\nMCsGA1UdIwQkMCKAIFzlB79h+DHFW5mqyY/1YTkZg2rqJjOCZHUpCR4smnA5MAoG\nCCqGSM49BAMCA0gAMEUCIQDRN6EeTBMeXgeCUlAu9sNYF9ILl8lRGuHR7BmX33fR\n7wIgbj2X7l2o8uCbhI2VvdsX86ZUpEfpI18/NwCrxa35zBw=\n-----END CERTIFICATE-----\n\x12H0F\x02!\x00\xd56\x9d\xa1\x89\xd8Y\x83\x1b\xad,gg\xe8\xb7\xef.\xd8\xf5Vg\xcf\xfe\x99-\x19\xefW9\xf6\xc4\x1a\x02!\x00\xfez\xa8\x8e1Ϥ6\\ݶ\xcd\xe1\x8b>VvV\xba.\xa8\xcc9:\xabl\xbc\"\xb1e\xf0\x98"

func main() {

	blockbytes := []byte(blockstring)
	a := common.Payload{}
	proto.Unmarshal(blockbytes, &a)
	fmt.Println(string(a.Header.ChannelHeader))

	channelHeader := common.ChannelHeader{}
	proto.Unmarshal(a.Header.ChannelHeader, &channelHeader)
	fmt.Println(channelHeader.TxId)

	b := peer.Transaction{}
	proto.Unmarshal(a.Data, &b)
	//fmt.Println(a.Header)
	//fmt.Println(b.Actions)
	for _, action := range b.Actions {
		tmpAction := peer.ChaincodeActionPayload{}
		proto.Unmarshal(action.Payload, &tmpAction)
		//fmt.Println(tmpAction.Action,tmpAction.ChaincodeProposalPayload)
		c := peer.ChaincodeProposalPayload{}
		proto.Unmarshal(tmpAction.ChaincodeProposalPayload, &c)
		//获取到输入参数
		//fmt.Println(string(c.Input))

		d := peer.ProposalResponsePayload{}
		proto.Unmarshal(tmpAction.Action.ProposalResponsePayload, &d)
		//fmt.Println(string(d.Extension))
		e := peer.ChaincodeAction{}
		proto.Unmarshal(d.Extension, &e)
		//fmt.Println(string(e.Events))

		f := pb.ChaincodeEvent{}
		proto.Unmarshal(e.Events, &f)
		//获取到合约事件
		fmt.Println(f.EventName, string(f.Payload))
	}
}
