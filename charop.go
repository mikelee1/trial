package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main()  {
	var str1 = "六年级数学天天练试题及答案<br/><br/>【题目】计数问题"
	fmt.Println(strings.Split(str1, " "))
	fmt.Println(strings.Fields(str1))
	aFunc := func(a rune) bool { return !unicode.IsLetter(a) }
	res := strings.FieldsFunc(str1, aFunc)
	fmt.Println(res[0][:9],res[len(res)-1])
}

