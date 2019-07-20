package main

import (
	"net/http"
	_ "net/http/pprof"
)

func foo() []byte {
	var buf [1000]byte
	return buf[:10]
}

var c int

func bar(b []byte) {
	c++
	for i := 0; i < len(b); i++ {
		b[i] = byte(c*i*i*i + 4*c*i*i + 8*c*i + 12*c)
	}
}

func main() {
	go http.ListenAndServe(":8200", nil)
	for {
		b := foo()
		bar(b)
	}
}
