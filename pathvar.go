package main

import (
	"fmt"
	"github.com/pathvar"
)

const (
	defaultConfigPath = "${GOPATH}/src/github.com/hyperledger/fabric/sdk/config/config.yaml"
)

func main() {
	path := pathvar.Subst(defaultConfigPath)
	fmt.Println(path)
}
