package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var GlobalDBClient *gorm.DB

//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "yunphant"
//	password = "toor"
//	dbname   = "sdzyb"
//)
const (
	//host     = "192.168.9.87"
	host     = "localhost"
	port     = 18255
	user     = "root"
	password = "YunPhant888"
	dbname   = "sdzyb"
)

type User struct {
	Name string
}

func init() {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		user, password, host, port, dbname)

	db, err := gorm.Open("mysql", mysqlInfo)

	if err != nil {
		db.Close()
		panic(err.Error())
	}

	GlobalDBClient = db

}

func main() {
	users := []*User{}
	if err := GlobalDBClient.Debug().Raw("select * from user").Find(&users).Error; err != nil {
		panic(err)
	}
	fmt.Println("users: ", users[0].Name)
}
