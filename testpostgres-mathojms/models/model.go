package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Openid            string    `gorm:"index"`
	Userid            string    `gorm:"index"` //用time.now.nano生成
	Nickname          string    `gorm:"default:'anonymous'"`
	Realname          string    //真实姓名
	Role              string    `gorm:"default:'student'"`
	Avatar            string    `gorm:"default:'null'"`
	Punish            string    `gorm:"default:'0'"`
	//Coin              int32     `gorm:"default:'500'"`
	Latitude          float64   `gorm:"default:'0'"` //纬度
	Longitude         float64   `gorm:"default:'0'"` //经度
	Geohash           string
	Onlycareselfgrade bool      `gorm:"default:false"`
	School            int //机构或者学校
}

func (t *User) TableName() string {
	return "mathojms_user"
}