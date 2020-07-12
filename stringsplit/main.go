package main

import (
	"strings"
	"fmt"
)

func main() {
	a := "aaa__bbb___cc"
	b := strings.Split(a, "_")
	fmt.Println(b)

	b = strings.SplitN(a, "_", 4)
	fmt.Println(b)
}
