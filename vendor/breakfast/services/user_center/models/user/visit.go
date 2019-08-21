package user

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Visit struct {
	gorm.Model
	Bevisitor int32 `gorm:"index"`
	Visitor   int32
	Visittime time.Time
	Inuse     string `gorm:"default:'1'"`
}

func (*Visit) TableName() string {
	return "breakfast_visit"
}
