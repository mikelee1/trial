package main

import (
	"fmt"
	"github.com/op/go-logging"
	"myproj.lee/try/testredis/config"
	"myproj.lee/try/testredis/core/redis"
	"os"
	"github.com/go-session/gin-session"
	"github.com/go-session/session"
	redis2 "github.com/go-session/redis"
	"context"
)

func init() {
	stdoutBackend := logging.NewBackendFormatter(
		logging.NewLogBackend(os.Stdout, "", 0),
		logging.MustStringFormatter(`%{color}[%{time:2006-01-02 15:04:05.000}] [%{module}] <%{shortfile}> %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`),
	)
	logging.SetBackend(stdoutBackend)
}

func main() {
	config.Init("", "./testredis/config/config.yaml")
	//config.Init("","./config/config.yaml")

	fmt.Println(config.GetConfig().Name)
	fmt.Println(config.GetConfig().Info.Id)
	fmt.Println(config.GetConfig().QabircConfig.Redis.Address)

	fmt.Println(config.GetConfig())
	redis.Init()
	redis.GetClient().Get("1001")
	store := redis2.NewRedisStore(&redis2.Options{
		Addr: config.GetConfig().QabircConfig.Redis.Address,
		DB:   config.GetConfig().QabircConfig.Redis.Db,
	})
	ginsession.New(
		session.SetStore(store),
		//设置session的有效时间，1小时
		session.SetExpired(3600),
	)
	//a := store.GetKeyValue(context.TODO(),"1001")
	//fmt.Println(a)
	b := store.GetKeyValue(context.TODO(),"aac4857f-59ed-4fe4-b770-3c5df6994dbd")
	fmt.Println(b)
	err := store.Delete(context.TODO(),"aac4857f-59ed-4fe4-b770-3c5df6994dbd")
	fmt.Println(err)



}
