package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	res := []byte{}
	for scanner.Scan() {

		str := scanner.Bytes()
		length := len(str)
		if length == 0 {
			break
		}
		res = []byte{}

		for k, _ := range str {
			res = append(res, str[length-1-k])
		}

	}

	fmt.Println(string(res))
}
