package main

import (
	"fmt"
	//"github.com/mikemintang/go-curl"
)

func main() {

	url := "http://192.168.9.82:6984"

	headers := map[string]string{
		"User-Agent":    "Sublime",
		"Authorization": "Bearer access_token",
		"Content-Type":  "application/json",
	}


	// 链式操作
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		Get()

	if err != nil {
		fmt.Println(err)
	} else {
		if resp.IsOk() {
			fmt.Println(resp.Body)
		} else {
			fmt.Println(resp.Raw)
		}
	}

}

//{"couchdb":"Welcome","version":"2.3.1","git_sha":"c298091a4","uuid":"d6b17088f8925a3d38e8a6c700468f77","features":["pluggable-storage-engines","scheduler"],"vendor":{"name":"The Apache Software Foundation"}}
//{"couchdb":"Welcome","version":"2.3.1","git_sha":"c298091a4","uuid":"686b30241a56a3cc7c7ce6ce999f654a","features":["pluggable-storage-engines","scheduler"],"vendor":{"name":"The Apache Software Foundation"}}