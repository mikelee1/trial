package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/redis"
	"github.com/go-session/gin-session"
	"github.com/go-session/session"
	"net/http"
	"fmt"
)

func main() {
	g := gin.Default()

	store := redis.NewRedisStore(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   1,
	})


	g.Use(
		ginsession.New(
			session.SetStore(store),
			//session.SetExpired(60),
		),
	)

	g.GET("/login", func(ctx *gin.Context) {
		//第一步 检查自己的sessionid在redis中是否存在，存在表示已经登录
		//第二步 若自己的sessionid在redis中不存在，则发送手机密码到后台取token（手机密码hash值），根据token查看是否已在别处登录
		//第三步 若已在别处登录，提示错误，否则，跳转首页并在redis中加入token:true

		//store is session.store, it is an interface
		store := ginsession.FromContext(ctx)
		//store is session.Manager, it is an interface, to communicate with
		ms,err := ginsession.GetManage(ctx)
		if err == nil{
			//mike 检测自己的sessionid是否存在，仅对本浏览器有效。
			fmt.Println("this is sessionid:",store.SessionID())
			fmt.Println("this is sessionid1:",store.SessionID())
			exist,_ := ms.GetOpts().GetStore().Check(ctx,store.SessionID())

			fmt.Println(exist)
			if exist {
				//ctx.AbortWithStatus(404)
				ctx.String(http.StatusOK,"error:already login")
				return
			}else{
				exist,_ := ms.GetOpts().GetStore().Check(ctx,"110")
				if exist{
					ctx.String(http.StatusOK,"error:another already login")
					return
				}
			}
		}else {
			panic(err)
		}

		store.Set("foo", "bar22")
		err = store.SaveCustom("110", "bar22")
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}
		err = store.Save()
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		ctx.Redirect(302, "/foo")
	})

	g.GET("/foo", func(ctx *gin.Context) {
		store := ginsession.FromContext(ctx)
		foo, ok := store.Get("foo")
		if !ok {
			//ctx.AbortWithStatus(404)
			ctx.String(http.StatusOK,"error:cant get foo")
			return
		}
		ctx.String(http.StatusOK, "foo:%s", foo)
	})

	g.GET("/logout", func(ctx *gin.Context) {
		store := ginsession.FromContext(ctx)
		err := store.DelCustom(store.SessionID())
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}
		err = store.DelCustom("110")
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}
		ctx.String(http.StatusOK, "logout ok")
	})

	g.Run(":8011")
}