package main

import (
	"os/exec"
	"fmt"
	"io/ioutil"
)

func Execcmd(cmdstring string) ([]byte) {
	cmd := exec.Command("/bin/ash", "-c", cmdstring)

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return nil
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return nil
	}

	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return nil
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return nil
	}


	return bytes
}