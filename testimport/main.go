package main

import (
	"fmt"
	_ "myproj/try/testimport/pkg1"
	"time"
)

func main() {
	fmt.Println("start")
	fmt.Println(time.Now())
	b := Inc()
	fmt.Println(b)
	fmt.Println(&b)
}

func Inc() (v int) {
	fmt.Println(&v)
	defer func() {
		v++
		fmt.Println(&v)
	}()
	return 42
}
