package ecies

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	curve := elliptic.P256()
	priv, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Printf("the generate key pair err : %v\n", err)
	}
	msg := []byte("helloworld")
	encrypt, err := EciesEncrypt(&(priv.PublicKey), msg, false)
	if err != nil {
		fmt.Printf("the encrypt err :%v\n", err)
	}
	decrypt, err := EciesDecrypt(priv, encrypt, false)
	if err != nil {
		fmt.Printf("the decrypt err :%v\n", err)
	}
	fmt.Println(string(decrypt))
}

func TestEncryptAndDecryptGM(t *testing.T) {
	curve := elliptic.P256()
	priv, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Printf("the generate key pair err : %v\n", err)
	}
	msg := []byte("helloworld")
	encrypt, err := EciesEncrypt(&(priv.PublicKey), msg, true)
	if err != nil {
		fmt.Printf("the encrypt err :%v\n", err)
	}
	decrypt, err := EciesDecrypt(priv, encrypt, true)
	if err != nil {
		fmt.Printf("the decrypt err :%v\n", err)
	}
	fmt.Println(string(decrypt))
}
