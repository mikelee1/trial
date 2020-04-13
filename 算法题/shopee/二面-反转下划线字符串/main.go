package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(transverse(input))
}

var input = "I__work__at__shopee"

func transverse(s string) string {
	var res []string
	//words := strings.Split(s, "_")
	var words []string
	j := 0
	flag := false
	for i := 0; i <= len(s)-1; i++ {
		if s[i] == byte('_') {
			flag = true
			continue
		}
		if flag {
			words = append(words, s[j:i])
			j = i
			flag = false
			continue
		}

	}
	words = append(words, s[j:])

	for _, word := range words {
		res = append(res, handle(word))
	}
	return strings.Join(res, "")

}

func handle(str string) string {
	res := ""
	tail := len(str) - 1
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == byte('_') {
			tail--
			continue
		}
		res = res + string(str[i])
	}
	res = res + string(str[tail+1:])
	return res
}
