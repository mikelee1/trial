package main

import (
	"os"
	"fmt"
)

func main() {
	src, err := os.Open("111.txt")
	defer src.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
