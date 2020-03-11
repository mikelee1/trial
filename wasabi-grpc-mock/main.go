package main

import (
	"fabric/core/comm"
	"time"
	"fmt"
)

const (
	defaultTimeout = time.Second * 3
)

type Endpoint struct {
	Address  string
	Override string
	TLS      []byte
	Timeout  time.Duration
}

func main() {
	e := Endpoint{
		Address:  "47.115.132.208:31003",
		Timeout:  time.Duration(3 * time.Second),
		Override: "orderer-2-baas3",
		TLS: []byte(`-----BEGIN CERTIFICATE-----
MIICVjCCAf2gAwIBAgIQM4cSyjj1If5qbjxewBm7KzAKBggqhkjOPQQDAjB2MQsw
CQYDVQQGEwJDTjERMA8GA1UECBMIWmhlamlhbmcxETAPBgNVBAcTCEhhbmd6aG91
MRAwDgYDVQQJEwdjb21wYW55MQ8wDQYDVQQREwYzMTAwMDAxDjAMBgNVBAoTBWJh
YXMzMQ4wDAYDVQQDEwViYWFzMzAeFw0yMDAzMTEwNTE1MDBaFw0zMDAzMDkwNTE1
MDBaMHYxCzAJBgNVBAYTAkNOMREwDwYDVQQIEwhaaGVqaWFuZzERMA8GA1UEBxMI
SGFuZ3pob3UxEDAOBgNVBAkTB2NvbXBhbnkxDzANBgNVBBETBjMxMDAwMDEOMAwG
A1UEChMFYmFhczMxDjAMBgNVBAMTBWJhYXMzMFkwEwYHKoZIzj0CAQYIKoZIzj0D
AQcDQgAEOrn2GiKcc0c7H5Cz6qxQSzvd/Njk+mCKpsTlPtsQ1FBBICGsZH1Gjejz
qbKGda6+bOCnBKEL9Twbd92eVeHiqaNtMGswDgYDVR0PAQH/BAQDAgGmMB0GA1Ud
JQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1Ud
DgQiBCC9mhtUPgLLrOWyTmbS2QGSfCHmpPk1p92IOZABTU1FLTAKBggqhkjOPQQD
AgNHADBEAiAcJo7xvpFVd4UMwvZkROMvaHdD8TUQHM/YFLtAZ03yBwIgIsBHzydj
2AxnRavDs087VL39rtpdLjubs78FsF3Tsmw=
-----END CERTIFICATE-----
`),
	}
	err := getClient(&e)
	if err != nil {
		fmt.Println(err)
	}
}

func getClient(endpoint *Endpoint) (error) {
	clientConfig := comm.ClientConfig{}
	timeout := endpoint.Timeout
	if timeout == time.Duration(0) {
		timeout = defaultTimeout
	}
	clientConfig.Timeout = timeout
	if endpoint.TLS != nil {
		secOpts := &comm.SecureOptions{
			UseTLS: true,
		}
		secOpts.ServerRootCAs = [][]byte{endpoint.TLS}
		clientConfig.SecOpts = secOpts
	}

	gClient, err := comm.NewGRPCClient(clientConfig)
	if err != nil {
		fmt.Println("Fail to new")
		return err
	}
	fmt.Println(endpoint)
	_, err = gClient.NewConnection(endpoint.Address, endpoint.Override)
	return err
}
