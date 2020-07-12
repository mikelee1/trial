package main_test

import (
	"testing"
	"myproj.lee/try/testpostgres-mathojms/models"
	"github.com/op/go-logging"
	logger2 "myproj.lee/try/common/logger"
	"github.com/jinzhu/gorm"
	"time"
	"github.com/gansidui/geohash"
	"math/rand"
	"strings"
	"strconv"
	"fmt"
)

var logger *logging.Logger
var dbClient *gorm.DB

func init() {
	logger = logger2.GetLogger()
	dbClient = models.InitDB()
	if dbClient == nil {
		logger.Error("fail")
		panic("fail")
	}
}

func Test_UpdateProblem(t *testing.T) {
	toreleasedid := []string{"1","2"}
	if err := dbClient.Debug().Model(&models.Problem{}).Where("id in (?)", toreleasedid).Update("status", "released").Error; err != nil {
		fmt.Println("err: ", err)
	}
}



func Test_Main(t *testing.T) {
	for i, _ := range [10]int{} {
		go SelectForUpdate(i)
	}
	time.Sleep(1 * time.Second)
}

func Test_SelectOne(t *testing.T) {
	tmp := &models.User{}
	if err := dbClient.Debug().Model(&models.User{}).Find(&tmp).Error; err != nil {
		logger.Info(err)
		return
	}
	logger.Info(tmp)
	time.Sleep(1 * time.Second)
}

func Test_InsertOne(t *testing.T) {
	tmp := &models.User{
		Userid:"2222",

	}
	if err := dbClient.Debug().Model(&models.User{}).Create(&tmp).Error; err != nil {
		logger.Info(err)
		return
	}
	logger.Info(tmp)
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

func Test_update(t *testing.T) {
	//以仓前为中心随机更新用户的经纬度
	longitude, latitude := 119.98954010009766, 30.28293228149414
	for k, _ := range [728]int{} {
		timens := int64(time.Now().Nanosecond())
		rand.Seed(timens)
		longitude += rand.Float64() - 0.5
		timens = int64(time.Now().Nanosecond())
		rand.Seed(timens)
		latitude += rand.Float64() - 0.5

		ns := geohash.GetNeighbors(latitude, longitude, 6)
		eighthashs := strings.Join(ns, "|")

		err := dbClient.Debug().Exec("update mathojms_user set latitude=?,longitude=?,geohash=?,eighthashs=? where id = ?",
			latitude, longitude, ns[0], eighthashs, k+1).Error
		if err != nil {
			logger.Error(err)
		}
		longitude, latitude = 119.98954010009766, 30.28293228149414

	}
}

func Test_updateuserid(t *testing.T) {
	for k, _ := range [728]int{} {
		userid := time.Now().UnixNano()
		useridStr := strconv.Itoa(int(userid))
		err := dbClient.Debug().Exec("update mathojms_user set userid=? where id = ?",
			useridStr, k+1).Error
		if err != nil {
			logger.Error(err)
		}
	}
}
