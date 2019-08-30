package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Jiaoan struct {
	gorm.Model
	Currenttime time.Time `gorm:"default:current_timestamp"`
	Title string
	Inuse string `gorm:"default:'1'"`
}

func (t *Jiaoan)TableName() string {
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