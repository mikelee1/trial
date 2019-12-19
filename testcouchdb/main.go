package main

import (
	"github.com/op/go-logging"
	"net/url"
	"github.com/zemirco/couchdb"
	"time"
)

var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("testcouchdb")
}

type couchDoc struct {
	ID           string
	Rev          string
	AccountMoney float64
	CreateTime   time.Time
	Name         string
}

func (c couchDoc) GetID() string {
	return c.ID
}

func (c couchDoc) GetRev() string {
	return c.Rev
}

// start
func main() {
	u, err := url.Parse("http://192.168.9.87:5984/")
	if err != nil {
		panic(err)
	}

	// create doc new client
	client, err := couchdb.NewAuthClient("admin", "pass", u)
	if err != nil {
		panic(err)
	}

	//获取couchdb的信息
	// get some information about your CouchDB
	//info, err := client.Info()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(info)

	//使用的数据库名
	db := client.Use("channel_sdzyb1")
	allDocs, err := db.AllDocs(&couchdb.QueryParameters{})
	for _, doc := range allDocs.Rows {
		cd := &couchDoc{}
		db.Get(cd, doc.ID)
		logger.Info(cd)
	}
}
