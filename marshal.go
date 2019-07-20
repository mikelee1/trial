package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type cc struct {
	Name string
	Info []struct{
		Status string
	}
}

func main()  {
	var ccone cc
	st1 := []string{`{
			"name":"A",
			"info":[
			{"status":"yes"},
			{"status":"no"}
			]}`}
	err := json.Unmarshal(convertStringsToBytes(st1),&ccone)
	//err := json.Unmarshal(str,&ccone)
	if err != nil{
		fmt.Printf("",err)
	}	else{
		fmt.Printf("god")
	}
	fmt.Println(ccone)
}


func convertStringsToBytes(stringContent []string) []byte {
	byteContent :=  strings.Join(stringContent, "")  // x20 = space and x00 = null
	return []byte(byteContent)
}