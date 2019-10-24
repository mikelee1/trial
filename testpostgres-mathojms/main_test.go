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

func Test_Main(t *testing.T) {
	for i, _ := range [10]int{} {
		go SelectForUpdate(i)
	}
	time.Sleep(1 * time.Second)
}

//查找jiaoan表
func SelectForUpdate(i int) {
	tmp := &models.User{}
	tx := dbClient.Begin()
	if err := tx.Debug().Set("gorm:query_option", "FOR UPDATE").Model(tmp).Where("id = ?", 2).Update("role", "mike").Error; err != nil {
		tx.Rollback()
		logger.Info(err)
		return
	}
	logger.Infof("%d got lock", i)
	tx.Commit()
	logger.Infof("%d release lock", i)
	logger.Info(tmp.Role)
}

func Test_Joins(t *testing.T) {
	tmp := []*models.User{}
	tx := dbClient.Begin()
	if err := tx.Debug().Table("mathojms_user u").Select("u.openid, u.userid").Joins("left join mathojms_school s on u.school = s.id").Scan(&tmp).Error; err != nil {
		tx.Rollback()
		logger.Info(err)
		return
	}
	tx.Commit()
	for _, v := range tmp {
		logger.Info(v.Openid, v.Userid)
	}
}

//对于存在的记录，for update对update和delete都有效
//对于不存在的记录，for update对update和delete无效¬
func Test_SelectForDelete(t *testing.T) {
	id := 1
	go SelectForDelete(id)
	go SelectForDelete1(id)
	time.Sleep(1 * time.Second)
}

func SelectForDelete(id int) {
	tmp := &models.User{}
	tx := dbClient.Begin()
	if err := tx.Debug().Set("gorm:query_option", "FOR UPDATE").Model(tmp).Where("id = ?", id).Update("role", "mike").Error; err != nil {
		tx.Rollback()
		logger.Info(err)
		return
	}
	logger.Infof("%d got lock", 1)
	tx.Commit()
	logger.Infof("%d release lock", 1)
	logger.Info(tmp.Role)
}

func SelectForDelete1(id int) {
	tmp := &models.User{}
	tx := dbClient.Begin()
	if err := tx.Debug().Set("gorm:query_option", "FOR UPDATE").Model(tmp).Where("id = ?", id).Delete(tmp).Error; err != nil {
		tx.Rollback()
		logger.Info(err)
		return
	}
	logger.Infof("%d got lock", 2)
	tx.Commit()
	logger.Infof("%d release lock", 2)
	logger.Info(tmp.Role)
}
