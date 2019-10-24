package models

import (
	"github.com/astaxie/beego/orm"
	"sync"
	"github.com/jinzhu/gorm"
	"time"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/op/go-logging"
	logger2 "myproj/try/common/logger"
)

var (
	once *sync.Once
	Oconnect orm.Ormer
	dbname = "mathojms"
	dbuser = "yunphant"
	dbpasswd = "yunphant2018"
	dbip = "127.0.0.1"
	dbport = "5432"
)

var logger *logging.Logger
func init()  {
	once = &sync.Once{}
	logger = logger2.GetLogger()
}
var db *gorm.DB

func InitDB() *gorm.DB {
	db = AutoMigrate()
	return db
}

func CreateConn() *gorm.DB {
	var err error
	if db == nil {
		db, err = gorm.Open("postgres", "host="+dbip+" port="+dbport+" user="+dbuser+"" +
			" dbname="+dbname+" password="+dbpasswd+" sslmode=disable")
		if err != nil {
			return nil
		}
		db.DB().SetMaxOpenConns(0)
		db.DB().SetMaxIdleConns(0)
		db.DB().SetConnMaxLifetime(10 * time.Minute)
	}
	return db
}

func AutoMigrate() *gorm.DB {
	db := CreateConn()
	if err := db.Exec("set transaction isolation level serializable").AutoMigrate(
		&User{},
	).Error; err != nil {
		logger.Panicf("Error auto-migrating database : %s", err)
	}
	return db
}

