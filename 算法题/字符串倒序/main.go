package main

import (
	"fmt"
	"strings"
)

var str = " k"

func main() {

	var (
		strBytes = []byte(str)
		resBytes []byte
	)

	words := strings.Split(str, " ")
	if len(words) > 100 {
		fmt.Println("Invalid string! ")
		return
	}

	byteLength := len(strBytes)
	for index := range strBytes {
		resBytes = append(resBytes, strBytes[byteLength-1-index])
	}
	fmt.Println(string(resBytes))
}
