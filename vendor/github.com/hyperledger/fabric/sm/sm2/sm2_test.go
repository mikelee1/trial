/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sm2

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	_ "encoding/pem"
	"fmt"
	"github.com/hyperledger/fabric/bccsp/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSM2LoadPrivKey(t *testing.T) {
	k, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	assert.NoError(t, err)
	raw, err := x509.MarshalECPrivateKey(k)
	assert.NoError(t, err)
	//x509.ParseECPrivateKey(raw)
	block, err := utils.PrivateKeyToPEM(k, []byte(""))
	assert.NoError(t, err)
	fmt.Println(block, err)
	bb := block
	fmt.Println("..........", raw, len(raw), base64.StdEncoding.EncodeToString(raw))
	fmt.Println("..........", bb, len(bb), base64.StdEncoding.EncodeToString(bb))
	ecKey, err := LoadSM2PrivKeyFromBytes(bb)
	assert.NoError(t, err)
	fmt.Printf("the ecKey is %v---%v\n", ecKey, err)
}

func TestSM2SignAndVerify(t *testing.T) {
	k, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	assert.NoError(t, err)
	//raw,err := x509.MarshalECPrivateKey(k)
	//x509.ParseECPrivateKey(raw)
	block, err := utils.PrivateKeyToPEM(k, []byte(""))
	assert.NoError(t, err)
	bb := block
	ecKey, err := LoadSM2PrivKeyFromBytes(bb)
	assert.NoError(t, err)
	fmt.Printf("the ecKey is %v---%v\n", ecKey, err)

	sign, err := SM2Sign(ecKey, []byte("helloworld"))
	assert.NoError(t, err)
	fmt.Println("the sign is ", sign, err)

	pub := k.PublicKey
	//fmt.Println("the pub is ",pub)
	pb, err := x509.MarshalPKIXPublicKey(&pub)
	assert.NoError(t, err)
	fmt.Println("the pb is ", pb, err)
	pEcKey, err := LoadSM2PubKeyFromByte(pb)
	//fmt.Println("the pEcKey is ",pEcKey)
	assert.NoError(t, err)
	verify, err := SM2Verify(pEcKey, []byte("helloworld"), sign)
	assert.NoError(t, err)
	assert.True(t, verify)
	fmt.Println("the verify is ", verify, err)

}
