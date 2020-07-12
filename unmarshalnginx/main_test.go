package main_test

import (
	"testing"
	"fmt"
)

type Inventory struct {
	Address string
}

func Test_Main(t *testing.T) {

	//sweaters := Inventory{
	//	Address: "192.168.9.82:30020",
	//}
	//tmpl, err := template.ParseFiles("nginx.conf")
	//w := &bytes.Buffer{}
	//CheckErr(err)
	//err = tmpl.Execute(w, sweaters)
	//fmt.Println(w.String())
	//CheckErr(err)

	a := fmt.Sprintf(`user  nginx;
worker_processes  1;

error_log  /var/log/nginx/error.log info;
pid        /var/run/nginx.pid;


events {
    worker_connections  1024;
}

stream {

    server {
        listen 7050;
        proxy_pass %s;
    }
}
`, "kjkjl")
	fmt.Printf(a)
}
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
