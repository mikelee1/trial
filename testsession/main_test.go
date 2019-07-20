package main_test

import (
	"net/http"

	"github.com/go-session/redis"
	"github.com/go-session/session"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"
)

func Test1(t *testing.T)  {

	app := gin.Default()

	store := redis.NewRedisStore(&redis.Options{
			Addr: "127.0.0.1:6379",
			DB:   1,
		})

	app.Use(
		ginsession.New(
			session.SetStore(store),
		),
	)

	app.GET("/", func(ctx *gin.Context) {
		store := ginsession.FromContext(ctx)
		store.Set("foo", "bar11")
		err := store.Save()
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		ctx.Redirect(302, "/foo")
	})

	app.GET("/foo", func(ctx *gin.Context) {
		store := ginsession.FromContext(ctx)
		foo, ok := store.Get("foo")
		if !ok {
			ctx.AbortWithStatus(404)
			return
		}
		ctx.String(http.StatusOK, "foo:%s", foo)
	})

	app.Run(":8011")

}