package main

import (
	"encoding/json"
	"fmt"
)

type AA struct {
	Value int
}

var aStr string

func init() {
	a := AA{Value: 1}
	aBytes, _ := json.Marshal(a)
	aStr = string(aBytes)
}

func main() {
	CorrectTest()
	ErrorTest()
}

func CorrectTest() {
	fmt.Println("---CorrectTest start---")

	//直接上链
	chainValue := []byte(aStr)
	aRes := string(chainValue)

	//sdk解析
	sdkHandle(aRes)

	fmt.Println("---CorrectTest end---")
}

func ErrorTest() {
	fmt.Println("---ErrorTest start---")
	//marshal后上链
	aMarshal, _ := json.Marshal(aStr)
	chainValue := aMarshal
	aRes := string(chainValue)

	//sdk解析
	sdkHandle(aRes)

	fmt.Println("---ErrorTest start---")
}

func sdkHandle(aRes string) {
	rawA := AA{}
	err := json.Unmarshal([]byte(aRes), &rawA)
	if err != nil {
		panic(err)
	}
}
