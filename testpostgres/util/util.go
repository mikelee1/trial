package util

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"time"
)

const (
	_host     = "localhost"
	_port     = "5432"
	_user     = "yunphant"
	_password = "yunphant2018"
	_dbname   = "breakfast"
)

var db *gorm.DB

// create db connection
func CreateConn() *gorm.DB {
	var err error
	if db == nil {
		db, err = gorm.Open("postgres", "host="+_host+" port="+_port+" user="+_user+" dbname="+_dbname+" password="+_password+" sslmode=disable")
		if err != nil {
			fmt.Println(err)
		}
		db.DB().SetMaxOpenConns(100)
		db.DB().SetMaxIdleConns(100)
		db.DB().SetConnMaxLifetime(10 * time.Minute)
	}
	return db
}

