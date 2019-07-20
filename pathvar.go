package main

import (
	"github.com/pathvar"
	"fmt"
)

const (
	defaultConfigPath = "${GOPATH}/src/github.com/hyperledger/fabric/sdk/config/config.yaml"
)

func main()  {
	path := pathvar.Subst(defaultConfigPath)
	fmt.Println(path)
}
