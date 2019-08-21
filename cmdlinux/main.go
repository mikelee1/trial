package main

import (
	"fmt"
	"myproj/try/common/execshell"
)

func main() {
	res, err := execshell.Run("ls cmdlinux && ls cmdlinux/dir1")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
