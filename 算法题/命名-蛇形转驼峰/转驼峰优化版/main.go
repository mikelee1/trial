// 转驼峰 优化版
package main

import (
	"fmt"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("(_|-)([a-zA-Z]+)")

func ToCamelCase(str string) string {
	camel := re.ReplaceAllString(str, " $2")
	camel = strings.Title(camel)
	camel = strings.Replace(camel, " ", "", -1)

	return camel
}

func main() {
	str := "the_start_boy"
	fmt.Println(ToCamelCase(str))
}