package main

import (
	"myproj/try/testredis/config"
	"fmt"
	"myproj/try/testredis/core/redis"
	"github.com/op/go-logging"
	"os"
)

func init()  {
	stdoutBackend := logging.NewBackendFormatter(
		logging.NewLogBackend(os.Stdout, "", 0),
		logging.MustStringFormatter(`%{color}[%{time:2006-01-02 15:04:05.000}] [%{module}] <%{shortfile}> %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`),
	)
	logging.SetBackend(stdoutBackend)
}


func main()  {
	config.Init("","./testredis/config/config.yaml")
	//config.Init("","./config/config.yaml")

	fmt.Println(config.GetConfig().Name)
	fmt.Println(config.GetConfig().Info.Id)
	fmt.Println(config.GetConfig().QabircConfig.Redis.Address)

	fmt.Println(config.GetConfig())
	redis.Init()

}