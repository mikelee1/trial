package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res []string
	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			break
		}
		res = append(res, Trans(str)...)
	}

	for _, v := range res {
		fmt.Println(v)
	}
}

func Trans(s string) []string {
	res := []string{}
	sBytes := []byte(s)
	remain := len(sBytes) % 8
	b := len(sBytes) / 8
	if remain != 0 {
		b = (len(sBytes) + (8 - remain)) / 8
		for j := 0; j < (8 - remain); j++ {
			sBytes = append(sBytes, []byte("0")...)
		}
	}

	for i := 0; i < b; i++ {
		res = append(res, string(sBytes[i*8:i*8+8]))
	}
	return res

}
