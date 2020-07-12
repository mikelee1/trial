package main

import (
	_ "myproj.lee/try/testinit/pkg"
	"time"
)

func init() {
	panic("")
}

func main() {
	time.Sleep(time.Second*4)
}
