//package main
//
//import (
//	"github.com/gin-gonic/gin"
//	"fmt"
//	"time"
//	_ "./docs"
//	//"github.com/swaggo/gin-swagger/swaggerFiles"
//	"github.com/swaggo/files"
//	"github.com/swaggo/gin-swagger"
//)
//
//func main() {
//	var err error
//	g := gin.New()
//	g.GET("/hello", func(context *gin.Context) {
//		time.Sleep(7*time.Second)
//		fmt.Println("hello")
//	})
//	g.GET("/mike", func(context *gin.Context) {
//		time.Sleep(5*time.Second)
//		fmt.Println("mike")
//	})
//	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
//	//g.Static("/swagger", "./testgin/swagger")
//	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,url))
//
//	//err := gen.New().Build("./","./testgin/main.go","./testgin/swagger","camelcase")
//	//if err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//	err = g.Run(":8080")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}


/**
 * Created by martin on 01/02/2019
 */

package main

import (
	_ "./docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

// @title 测试swagger
// @version 1.0.0
// @description  测试swagger
// @BasePath /api/v1/
func main() {
	r := gin.New()

	// 创建路由组
	v1 := r.Group("/api/v1")

	v1.GET("/record/:userId",record)
	v1.GET("/hello/:name", hello)

	// 文档界面访问URL
	// http://127.0.0.1:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

// @获取指定ID记录
// @Description get record by ID
// @Accept  json
// @Produce json
// @Param   some_id     path    int     true        "userId"
// @Success 200 {string} string	"ok"
// @Router /record/{some_id} [get]
func record(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// @欢迎
// @Description greet
// @Accept  json
// @Produce json
// @Param   name     path    string     true        "name"
// @Success 200 {string} string	"ok"
// @Router /hello/{name} [get]
func hello(c *gin.Context) {
	time.Sleep(2*time.Second)
	c.String(http.StatusOK,"hello,"+c.Param("name"))
}