package main

//import (
//	"github.com/gorilla/websocket"
//	"net/http"
//	"fmt"
//	"github.com/gin-gonic/gin"
//)

//import (
//	"github.com/gin-gonic/gin"
//	"myproj.lee/try/gin_websocket/controller"
//	"github.com/gorilla/websocket"
//	"net/http"
//	"fmt"
//)
//
//func main()  {
//	router := gin.Default()
//
//	router.LoadHTMLGlob("/Users/leemike/go/src/myproj.lee/try/gin_websocket/templates/*")
//
//	router.GET("/home",controller.GetHome)
//	router.GET("/", controller.GetIndex)
//	//websocket
//
//
//
//
//	router.Run("0.0.0.0:8089")
//}

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, []byte("pong"))
	}
}

func main() {

	r := gin.Default()

	//websocket 请求使用 wshandler函数处理
	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})
	r.LoadHTMLGlob("/Users/leemike/go/src/myproj.lee/try/gin_websocket/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.Run("0.0.0.0:8089")
}
