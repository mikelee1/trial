package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(transverse("i love shopee !"))
}

func transverse(s string) string {
	var res []string
	words := strings.Split(s, " ")
	for _, word := range words {
		res = append(res, handle(word))
	}
	return strings.Join(res, " ")
}

func handle(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res = res + string(s[i])
	}
	return res
}
