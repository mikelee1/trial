package main

import (
	"bytes"
	"github.com/go-ego/ego/mid/json"
	"fmt"
)

func main() {
	bufferdata := struct {
		Name string
	}{"mike"}
	bufferBytes,_ := json.Marshal(bufferdata)
	a := bytes.NewBuffer(bufferBytes)
	fmt.Println(a.String())
}
