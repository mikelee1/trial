package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		strList := strings.Split(line, " ")
		res := strList[len(strList)-1]
		fmt.Println(len([]byte(res)))
	}
}
