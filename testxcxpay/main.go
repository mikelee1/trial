package main

import (
	"github.com/medivhzhan/weapp/token"
	"fmt"
)

const (
	appID = "wx69ccfde6a7c19a05"
	secret = "51b16a4dfc13cc58338db39832ee3b7a"
)

func main() {
	// 获取次数有限制 获取后请缓存
	tok, exp, err := token.AccessToken(appID, secret)
	fmt.Println(tok,exp,err)
}
