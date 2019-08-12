package sm4

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSM4EncryptDecrypt(t *testing.T) {
	key, err := GetRandomBytes(16)
	assert.NoError(t, err)
	fmt.Println("the key is ", key)
	src := []byte("h")
	fmt.Println("the src is ", src)
	encrypt, err := Encrypt(key, src)
	assert.NoError(t, err)
	fmt.Println("the encrypt src is ", encrypt)
	decrypt, err := Decrypt(key, encrypt)
	assert.NoError(t, err)
	fmt.Println("the decrypt string is ", string(decrypt))
	assert.Equal(t, src, decrypt)
}
