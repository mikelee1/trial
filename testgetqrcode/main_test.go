package main_test

import (
	"testing"
	"fmt"
	"os"
)

func Test_main(t *testing.T) {
	fi, err := os.Stat("./lee.jpg")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Println(fi.Name())
}
