package main

import (
	"fmt"
	"myproj.lee/try/common/execshell"
)

func main() {
	res, err := execshell.Run("ls cmdlinux && ls cmdlinux/dir1")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	res, err = execshell.Run("cd cmdlinux && tar zcvf dir1.tar.gz dir1")
	if err != nil{
		fmt.Println(err)
		return
	}
}
