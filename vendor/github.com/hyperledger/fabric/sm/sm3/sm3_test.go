package sm3

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	h := New()
	h.Write([]byte("xxxx"))
	r := h.Sum(nil)
	fmt.Println(r)
	assert.Equal(t, r, []byte{206, 242, 50, 24, 76, 249, 163, 158, 201, 129, 70, 111, 85, 233, 254, 49, 229, 39, 62, 221, 189, 19, 245, 97, 170, 12, 54, 204, 136, 197, 144, 230})
}
