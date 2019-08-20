package ssh

import (
	"github.com/appleboy/easyssh-proxy"
	"time"
	"fmt"
)

func NewSshClient(server string) *easyssh.MakeConfig {
	ssh := &easyssh.MakeConfig{
		User:   "centos",
		Server: server,
		KeyPath: "/Users/leemike/.ssh/id_rsa",
		Port:    "22",
		Timeout: 60 * time.Second,
	}

	// Call Run method with command you want to run on remote server.
	_, _, _, err := ssh.Run("cd /home/centos/go/src/wasabi && pwd && ls", 60*time.Second)
	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
		return nil
	} else {
		fmt.Println("connect ssh "+server+" ok")
	}
	return ssh
}
