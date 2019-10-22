package main_test

import (
	"testing"
	"fmt"
)

var GlobalTemp *Temp

type Temp struct {
	Name string
}

//打印出值修改后的值
func Test_Value(t *testing.T) {
	var temp = Temp{
		Name: "aaa",
	}
	defer fmt.Println(temp)
	temp.Name = "bbb"
	return
}

//打印出指针修改后的值
func Test_Point(t *testing.T) {
	var temp = &Temp{
		Name: "aaa",
	}
	defer fmt.Println(temp)
	temp.Name = "bbb"
	return
}

//打印出全局的值
func Test_GlobalValue(t *testing.T) {

	defer fmt.Println(GlobalTemp)
	GlobalTemp = &Temp{
		Name: "aaa",
	}
	GlobalTemp.Name = "bbb"
	return
}

//打印出全局的值
func Test_GlobalValuePost(t *testing.T) {
	GlobalTemp = &Temp{
		Name: "aaa",
	}
	GlobalTemp.Name = "bbb"
	defer fmt.Println(GlobalTemp)
	return
}

//直接打印err
func Test_NoFunc(t *testing.T) {
	var err error
	defer fmt.Println(err)
	err = fmt.Errorf("skdfj")
	return
}

//打印出err改变后的值
func Test_WithFunc(t *testing.T) {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = fmt.Errorf("skdfj")
	return
}
