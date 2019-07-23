package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

func main() {

	u := User{}
	for k, _ := range [10]int{} {
		go func(v int) {
			db := CreateConn()
			if err := db.Raw("select * from breakfast_user").Scan(&u).Error; err != nil {
				fmt.Println(err)
			}
			fmt.Println(v, u)
		}(k)
	}

	time.Sleep(2 * time.Second)
}

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type User struct {
	Openid   string
	Nickname string
	Avatar   string
	Punish   string
	Inuse    string
	Phone    string
}

var db *gorm.DB

func CreateConn() *gorm.DB {

	var dbConfig = &DBConfig{
		Host:     "127.0.0.1",
		Port:     "5432",
		Name:     "breakfast",
		User:     "yunphant",
		Password: "yunphant2018",
	}
	var err error
	if db == nil {
		db, err = gorm.Open("postgres", "host="+dbConfig.Host+" port="+dbConfig.Port+" user="+dbConfig.User+" dbname="+dbConfig.Name+" password="+dbConfig.Password+" sslmode=disable")
		if err != nil {
			fmt.Println(err)
			return nil
		}
		db.DB().SetMaxOpenConns(0)
		db.DB().SetMaxIdleConns(0)
		db.DB().SetConnMaxLifetime(10 * time.Minute)

	}
	return db
}
