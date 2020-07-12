package main_test

import (
	"testing"
	"fmt"
	"errors"
)

type T struct {
	Age      int
	Name     string
	Children []int
}

func Test_main(t1 *testing.T) {
	// 初始化测试用例

	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()
	err = errors.New("sldkjflk")
}
