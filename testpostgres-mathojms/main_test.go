package main_test

import (
	"testing"
	"myproj/try/testpostgres-mathojms/models"
	"github.com/op/go-logging"
	logger2 "myproj/try/common/logger"
	"github.com/jinzhu/gorm"
	"time"
)

var logger *logging.Logger
var dbClient *gorm.DB

func init() {
	logger = logger2.GetLogger()
	dbClient = models.InitDB()
	if dbClient == nil {
		logger.Error("fail")
	}
}

func Test_Main(t *testing.T)  {
	for i, _ := range [10]int{}{
		go SelectForUpdate(i)
	}
	time.Sleep(1*time.Second)
}

//查找jiaoan表
func SelectForUpdate(i int) {
	tmp := &models.User{}
	tx := dbClient.Begin()
	if err := tx.Set("gorm:query_option", "FOR UPDATE").Model(&models.User{}).Find(tmp).Error; err != nil {
		tx.Rollback()
		logger.Info(err)
		return
	}
	logger.Infof("%d got lock", i)
	if err := tx.Model(&models.User{}).Update("role", "mike").Error; err != nil {
		tx.Rollback()
		logger.Info(err)
		return
	}
	tx.Commit()
	logger.Infof("%d release lock", i)
	logger.Info(tmp.Role)
}
