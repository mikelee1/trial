package main

import "fmt"

func main() {
	fmt.Println(subString("hello", "ll"))
}

func subString(origin string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	needleLength := len(needle)
	for i := 0; i <= len(origin)-needleLength; i++ {
		if origin[i:i+needleLength] == needle {
			return i
		}
	}
	return -1
}
