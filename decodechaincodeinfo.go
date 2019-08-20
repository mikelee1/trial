package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type PeerChaincodeStatus struct {
	NodeName   string
	NodeStatus string
}

func main() {
	var nccs []*PeerChaincodeStatus
	pdata, _ := base64.StdEncoding.DecodeString("W3siTm9kZU5hbWUiOiJwZWVyLTEtYmFhczQiLCJOb2RlU3RhdHVzIjoiaW5zdGFudGlhdGVkIn0seyJOb2RlTmFtZSI6InBlZXItMC1iYWFzNCIsIk5vZGVTdGF0dXMiOiJ1cGxvYWRlZCJ9XQ==")
	err := json.Unmarshal(pdata, &nccs)
	if err != nil {
		fmt.Println(err)
		return
	}
	for key, value := range nccs {
		fmt.Println(key,value)
	}
}
