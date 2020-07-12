package main

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
)

func main() {
	decrypto([]byte(`-----BEGIN CERTIFICATE-----
MIICLTCCAdOgAwIBAgIQQtysHGJgEVMVSUjHy+MbyDAKBggqhkjOPQQDAjB2MQsw
CQYDVQQGEwJDTjERMA8GA1UECBMIWmhlamlhbmcxETAPBgNVBAcTCEhhbmd6aG91
MRAwDgYDVQQJEwdjb21wYW55MQ8wDQYDVQQREwYzMTAwMDAxDjAMBgNVBAoTBWJh
YXMxMQ4wDAYDVQQDEwViYWFzMTAeFw0yMDA0MjUxMTE4MDBaFw0zMDA0MjMxMTE4
MDBaMGwxCzAJBgNVBAYTAkNOMREwDwYDVQQIEwhaaGVqaWFuZzERMA8GA1UEBxMI
SGFuZ3pob3UxEDAOBgNVBAkTB2NvbXBhbnkxDzANBgNVBBETBjMxMDAwMDEUMBIG
A1UEAwwLQWRtaW5AYmFhczEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASil1BR
5eWhBuWyFkdkbsg666oQdsDcwouHcwuGyg+Z5je4U8uw5VnKbfdrNzD5FqHw7JW3
Afocs6WlDv0Bi0Yqo00wSzAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/BAIwADAr
BgNVHSMEJDAigCAGwL+k5pBk5cGHA6zXBiRn6J95BhgXK5UWt//JU9pLPTAKBggq
hkjOPQQDAgNIADBFAiEAo0mZulfOXSE3tpw1y34W1+wpbZ8+o0bzFPu8zLN0Ke8C
IFqL4PTcUh15gk635lc2j5LNsD6z4wv9e2v6NL+DH1d2
-----END CERTIFICATE-----`))
}

func decrypto(creatorByte []byte) error {
	//fmt.Println("creator: ", string(creatorByte))

	certStart := bytes.IndexAny(creatorByte, "-----BEGIN")
	if certStart == -1 {
		fmt.Errorf("No certificate found")
		return fmt.Errorf("No certificate found")
	}
	//creatorOrg := string(creatorByte[2 : certStart-3])
	certText := creatorByte[certStart:]
	bl, _ := pem.Decode(certText)
	if bl == nil {
		fmt.Errorf("Could not decode the PEM structure")
		return fmt.Errorf("Could not decode the PEM structure")
	}

	cert, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		fmt.Errorf("ParseCertificate failed")
		return err
	}
	PrettyPrint(cert)
	//fmt.Println(cert.Subject.CommonName)
	//fmt.Println(cert.Issuer.Organization)
	return nil
}

func PrettyPrint(a interface{}) {
	ab, _ := json.MarshalIndent(a, "", "    ")
	fmt.Println(string(ab))
}
