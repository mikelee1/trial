package main

import (
	"fabric/core/comm"
	"fmt"
	"time"
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
		Address:  "192.168.9.201:7051",
		Timeout:  time.Duration(3 * time.Second),
		Override: "peer0.org1.example.com",
		//ca.crt文件
		TLS: []byte(`-----BEGIN CERTIFICATE-----
MIICVzCCAf6gAwIBAgIRAKL86iZxDvzlNj6oHvAVSIMwCgYIKoZIzj0EAwIwdjEL
MAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG
cmFuY2lzY28xGTAXBgNVBAoTEG9yZzEuZXhhbXBsZS5jb20xHzAdBgNVBAMTFnRs
c2NhLm9yZzEuZXhhbXBsZS5jb20wHhcNMjAwOTIzMDY1MDAwWhcNMzAwOTIxMDY1
MDAwWjB2MQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UE
BxMNU2FuIEZyYW5jaXNjbzEZMBcGA1UEChMQb3JnMS5leGFtcGxlLmNvbTEfMB0G
A1UEAxMWdGxzY2Eub3JnMS5leGFtcGxlLmNvbTBZMBMGByqGSM49AgEGCCqGSM49
AwEHA0IABAjbuBtIwlJSoDZZ47XQfj/xzUtQ1uTYhe0GrtpBeIJKepZc/p/MCRoh
JknoZsN2SIsARsc+VzH3SUQWaUVYvbCjbTBrMA4GA1UdDwEB/wQEAwIBpjAdBgNV
HSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwDwYDVR0TAQH/BAUwAwEB/zApBgNV
HQ4EIgQg5F4m43UdAwFumR3h3L+AMy5rCkHEaYUixaLzDfcfmWMwCgYIKoZIzj0E
AwIDRwAwRAIgaYhucnY7TVex07CysWpU7XEuNBWyhMvXVUzJ3/msIUMCIH/G8lO4
UyW3k69HDpsNoMBkrfvaE9f1DBJl1z2asS7G
-----END CERTIFICATE-----
`),
	}
	err := getClient(&e)
	if err != nil {
		fmt.Println(err)
	}
}

func getClient(endpoint *Endpoint) error {
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
	//fmt.Println(endpoint)
	_, err = gClient.NewConnection(endpoint.Address, endpoint.Override)
	return err
}
