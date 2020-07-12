package main


import (
	"fmt"
	"time"

	"github.com/appleboy/easyssh-proxy"
	"github.com/op/go-logging"
	"myproj.lee/try/common/logger"
)
var logger1 *logging.Logger

func init()  {
	logger1 = logger.GetLogger()
}

func main() {
	// Create MakeConfig instance with remote username, server address and path to private key.
	ssh := &easyssh.MakeConfig{
		User:   "centos",
		Server: "192.168.9.82",
		KeyPath: "/Users/leemike/.ssh/id_rsa",
		Port:    "22",
		Timeout: 60 * time.Second,
	}

	// Call Run method with command you want to run on remote server.
	stdout, _, _, err := ssh.Run("cd /home/centos/go/src/wasabi && pwd && ls", 60*time.Second)
	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Printf("stdout is :\n %s", stdout)
	}

}