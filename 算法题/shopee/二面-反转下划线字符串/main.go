package main

import (
	"fmt"
	"strings"
)

var input = "_I__work__at___yunphant_"

func main() {
	fmt.Println(transverseString1(input))
	fmt.Println(transverseString2(input))

}
//方法一: O(n)
func transverseString1(str string) string {
	var res []string
	var stack []string
	for _, s := range str {
		if string(s) == "_" {
			res = append(res, stack...)
			stack = []string{}
			res = append(res, "_")
			continue
		}
		stack = append([]string{string(s)}, stack...)
	}
	if len(stack) != 0 {
		res = append(res, stack...)
	}
	return strings.Join(res, "")
}

//方法二: O(n^2)
func transverseString2(s string) string {
	var res []string
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
		res = append(res, transverseWord(word))
	}
	return strings.Join(res, "")
}

func transverseWord(str string) string {
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
