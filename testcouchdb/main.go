package main

import (
	"github.com/op/go-logging"
	"net/url"
	"github.com/zemirco/couchdb"
	"fmt"
)

var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("testcouchdb")
}

type dummyDocument struct {
	couchdb.Document
	Foo  string `json:"foo"`
	Beep string `json:"beep"`
}

// start
func main() {
	u, err := url.Parse("http://192.168.9.87:5984/")
	if err != nil {
		panic(err)
	}

	// create a new client
	client, err := couchdb.NewAuthClient("admin", "pass", u)
	if err != nil {
		panic(err)
	}

	// get some information about your CouchDB
	info, err := client.Info()
	if err != nil {
		panic(err)
	}
	fmt.Println(info)
	db1 := client.Use("publicchain_publicchaincode")
	as, err := db1.AllDocs(&couchdb.QueryParameters{})

	for _, a := range as.Rows {
		fmt.Printf("id: %s, key: %v, value: %v, doc: %v \n", a.ID, a.Key, a.Value, a.Doc)
	}


}
