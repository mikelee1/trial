package main

import (
	"github.com/pingcap/tidb/config"
	"github.com/pingcap/tidb/store/tikv"
	"golang.org/x/net/trace"
)

func main() {
	_, _ = tikv.NewRawKVClient([]string{"127.0.0.1:2379"}, config.Security{})
	_ = trace.New("", "")
}
