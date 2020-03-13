package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	strs := ""
	target := ""
	count := 0
	for scanner.Scan() {
		if count == 0 {
			strs = scanner.Text()
		}
		if count == 1 {
			target = scanner.Text()
		}
		count++
		if count > 1 {
			break
		}
	}
	fmt.Println(strs, target)
	fmt.Println(byteCount(strs, target))
}

func byteCount(raw, target string) int {
	rawBytes := strings.Split(raw, "")
	count := 0
	upper := strings.ToUpper(target)
	for _, b := range rawBytes {

		if b == target || upper == strings.ToUpper(b) {
			count++
		}
	}
	return count
}
