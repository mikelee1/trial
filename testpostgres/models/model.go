package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/lib/pq"
)

type Jiaoan struct {
	gorm.Model
	Currenttime time.Time `gorm:"default:current_timestamp"`
	Title       string
	Inuse       string    `gorm:"default:'1'"`
}

func (t *Jiaoan) TableName() string {
	return "edu_jiaoan"
}

type PrincipalJiaoan struct {
	gorm.Model
	Jiaoan    int32
	Principal int32

	Inuse int32 `gorm:"default:1"`
}

func (*PrincipalJiaoan) TableName() string {
	return "edu_principal_jiaoan"
}

type TeacherAuth struct {
	gorm.Model
	Teacher int32
	Auth    pq.Int64Array `gorm:"type:integer[]"`

	Inuse int32 `gorm:"default:1"`
}

func (*TeacherAuth) TableName() string {
	return "teacher_auth"
}



type User struct {
	gorm.Model
	Openid            string    `gorm:"index"`
	Userid            string    `gorm:"index"` //用time.now.nano生成
	Nickname          string    `gorm:"default:'anonymous'"`
	Realname          string                               //真实姓名
	Role              string    `gorm:"default:'student'"` //有效值是student\teacher\pre-principal\principal，pre-principal是独立老师
	Avatar            string    `gorm:"default:'null'"`
	Punish            string    `gorm:"default:'0'"`
	Latitude          float64   `gorm:"default:'0'"` //纬度
	Longitude         float64   `gorm:"default:'0'"` //经度
	Geohash           string                         //自己经纬度所在的hash
	Eighthashs        string                         //自己经纬度周围八个区域的hash拼接
	Grade             GradeType `gorm:"default:4"`   //从二年级0开始
	Onlycareselfgrade bool      `gorm:"default:false"`
	Phone             string //手机号码
	Selfdesc          string //自我描述
}

func (*User) TableName() string {
	return "mathojms_user"
}

type GradeType int
