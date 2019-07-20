package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func GetHome(c *gin.Context)  {
	c.HTML(http.StatusOK,"home.html",gin.H{})
}


func GetIndex(c *gin.Context) {
	fmt.Printf("c.writer:%v,c.request:%v\n",c.Writer,c.Request)
	http.ServeFile(c.Writer, c.Request, "index.html")
}