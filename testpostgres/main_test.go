package main_test

import (
	"testing"
	"myproj.lee/try/testpostgres/models"
	"github.com/op/go-logging"
	logger2 "myproj.lee/try/common/logger"
	"github.com/jinzhu/gorm"
	"fmt"
	"strings"
	"jiaoan/services/admin_center/models/admin"
	"strconv"

	"github.com/gansidui/geohash"
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

//查找jiaoan表
func Test_getonejiaoan(t *testing.T) {
	tmp := &models.Jiaoan{}
	//如果jiaoan表里面已经有要create的条目，则会报"pq: duplicate key value violates unique constraint"
	if err := dbClient.Debug().Model(&models.Jiaoan{}).Where(&models.Jiaoan{Title: ""}).Where("id = ?", 1).Find(tmp).Error; err != nil {
		logger.Info(err)
		return
	}
	logger.Info(tmp)
}

//修改jiaoan表
func Test_getadmin(t *testing.T) {
	tmp := &admin.Administrator{}
	//判断record是否存在，需要用find
	logger.Info(dbClient.Debug().Model(&admin.Administrator{}).Where("phone = ?", "mike").Find(tmp).RecordNotFound())
	////如果jiaoan表里面已经有要create的条目，则会报"pq: duplicate key value violates unique constraint"
	if err := dbClient.Debug().Model(&admin.Administrator{}).Where("phone = ?", "mike").Find(tmp).Error; err != nil {
		logger.Info(err)
		return
	}
	logger.Info(tmp)
}

func Test_Addjiaoan(t *testing.T) {
	tmp := &models.Jiaoan{
		Title: "sdf1",
	}

	tx := dbClient.Begin()
	if err := tx.Debug().Model(&models.Jiaoan{}).Create(tmp).Error; err != nil {
		tx.Rollback()
		logger.Info(err)
		return
	}
	tx.Commit()

}

func Test_Addmultijiaoan(t *testing.T) {
	tmp := []*models.Jiaoan{
		&models.Jiaoan{
			Title: "sdf1",
		},
		&models.Jiaoan{
			Title: "sdf2",
		},
	}
	a := &models.Jiaoan{}
	sqlString := fmt.Sprintf("insert into %s (%s) values ", a.TableName(), "title")
	strs := []string{}
	for _, v := range tmp {
		strs = append(strs, fmt.Sprintf("('%s')", v.Title))
	}
	rawstr := strings.Join([]string{sqlString, strings.Join(strs, ",")}, " ")
	tx := dbClient.Begin()
	if err := tx.Debug().Model(&models.Jiaoan{}).Exec(rawstr).Error; err != nil {
		tx.Rollback()
		logger.Info(err)
		return
	}
	tx.Commit()

}

func Test_getoneprincipal(t *testing.T) {
	tmpjiaoans := []*models.PrincipalJiaoan{}
	jiaoanT := models.PrincipalJiaoan{}
	ids := []string{}
	for _, v := range []int{} {
		ids = append(ids, strconv.Itoa(v))
	}
	sqlstr1 := strings.Join(ids, ",")
	sqlstr := fmt.Sprintf("select id,jiaoan,principal from %s where principal in (%s)", jiaoanT.TableName(), sqlstr1)

	logger.Info(sqlstr)
	dbClient.Debug().Model(&models.Jiaoan{}).Exec(sqlstr).Find(&tmpjiaoans)
	logger.Info(tmpjiaoans)
}

func Test_AddTeacherAuth(t *testing.T) {
	tmp := &models.TeacherAuth{
		Teacher: 1,
		Auth:    []int64{1, 2},
	}

	tx := dbClient.Begin()
	if err := tx.Debug().Model(&models.TeacherAuth{}).Create(tmp).Error; err != nil {
		tx.Rollback()
		logger.Info(err)
		return
	}
	tx.Commit()

}

//location精度由6变为5以后，线上的定位数据需要相应更新
func Test_Update_Location(t *testing.T) {
	var err error
	logger.Info(err)
	users := []*models.User{}
	if err := dbClient.Model(&models.User{}).Find(&users).Error; err != nil {
		logger.Info(err)
		return
	}
	for _, u := range users {
		logger.Info(u.Nickname, u.Realname, u.Latitude, u.Longitude)
		////生成hash码
		hashs := geohash.GetNeighbors(float64(u.Latitude), float64(u.Longitude), 5)
		tmpu := models.User{
			Longitude:  float64(u.Longitude),
			Latitude:   float64(u.Latitude),
			Geohash:    hashs[0],
			Eighthashs: strings.Join(hashs[1:], "|"),
		}
		if err = dbClient.Debug().Model(&models.User{}).Where("id = ?", u.ID).Updates(tmpu).Error; err != nil {
			logger.Error(err)
			return
		}
	}

	logger.Info(users)
}
