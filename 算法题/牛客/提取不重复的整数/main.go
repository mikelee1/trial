package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	m := map[byte]bool{}
	for scanner.Scan() {
		res := []byte{}
		bytes := scanner.Bytes()
		for i := len(bytes) - 1; i >= 0; i-- {
			if ok := m[bytes[i]]; ok {
				continue
			}
			m[bytes[i]] = true
			res = append(res, bytes[i])
		}
		fmt.Println(string(res))
	}
}
