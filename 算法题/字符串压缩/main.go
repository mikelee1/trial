package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compressString("abbccd"))
}

func compressString(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}
	index := 0
	old := s[0]
	count := 0
	res := ""
	//same := false
	for index < length {
		if s[index] == old {
			//same = true
			index++
			count++
			continue
		}
		old = s[index]
		//same = false
		res = res + string(s[index-1]) + strconv.Itoa(count)
		count = 1
		index++
	}

	res = res + string(s[index-1]) + strconv.Itoa(count)

	if len(res) >= length {
		return s
	}

	return res
}
