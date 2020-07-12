package models

import (
	"github.com/jinzhu/gorm"
	"time"
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

}

func (t *User) TableName() string {
	return "mathojms_user"
}

type School struct {
	gorm.Model
	Name      string
	Principal int    `gorm:"index"` //属于哪个校长
	Address   string
	Geohash   string
	Inuse     string `gorm:"default:'1'"`
}

func (*School) TableName() string {
	return "mathojms_school"
}

type Problem struct {
	gorm.Model
	Asker       string    `column:"user__user_openid"`
	ProblemPic  string    `gorm:"default:''"`
	ProblemPic1 string    `gorm:"default:''"`
	ProblemPic2 string    `gorm:"default:''"`
	ProblemPic3 string    `gorm:"default:''"`
	Description string
	Asktime     time.Time `gorm:"default:current_timestamp"` //提问时间
	Grade       string    `gorm:"default:'二年级'"`             //问题的年级
	Category    string                                       //问题所属的知识点
	Jinxuan     string    `gorm:"default:'no'"`              //是否精选
	Inuse       string    `gorm:"default:'1'"`
	School      int       `gorm:"default:'0'"`        //问题属于哪个学校，方便学校进行汇总
	Status      string    `gorm:"default:'unserved'"` //状态：unserved\served\released 是否被本机构回答了
}

func (*Problem) TableName() string {
	return "mathojms_problem"
}