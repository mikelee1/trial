package main

import (
	"net/url"
	"fmt"
)

func main() {
	u,_ := url.Parse("http://192.168.9.11:2376")
	fmt.Println(u.Scheme)
	fmt.Println(u.Hostname())
	fmt.Println(fmt.Sprintf("%s://%s",u.Scheme,u.Hostname()))
}
