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

package gm

import (
	"fmt"
	"github.com/hyperledger/fabric/bccsp"
	_ "github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateSM4(t *testing.T) {
	b, err := New(nil)
	k, err := b.KeyGen(&bccsp.SM4KeyGenOpts{})
	fmt.Printf("the k and err is %v --->%v\n", k, err)
	cip, err := b.Encrypt(k, []byte("helloworld"), nil)
	fmt.Printf("the cip is %v and err is %v\n", cip, err)
	plain, err := b.Decrypt(k, cip, nil)
	fmt.Printf("the plain is %v and err is %v\n", string(plain), err)
}
func TestImportSM4(t *testing.T) {
	b, err := New(nil)
	k, err := b.KeyGen(&bccsp.SM4KeyGenOpts{})
	fmt.Printf("the k and err is %v --->%v\n", k, err)

	keyBytes, err := k.Bytes()
	k2, err := b.KeyImport(keyBytes, &bccsp.SM4ImportKeyOpts{})
	fmt.Printf("the k2 is %v and err is %v\n", k2, err)

	cip, err := b.Encrypt(k2, []byte("helloworld"), nil)
	fmt.Printf("the cip is %v and err is %v\n", cip, err)
	plain, err := b.Decrypt(k, cip, nil)
	fmt.Printf("the plain is %v and err is %v\n", string(plain), err)
}

func TestEncryptSM4(t *testing.T) {

}
func TestDecryptSM4(t *testing.T) {

}
