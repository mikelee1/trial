package user

import (
	e "breakfast/common/errors"
	"github.com/jinzhu/gorm"

	"breakfast/common/errors"
	"time"
)

type Subscribeuser struct {
	gorm.Model
	Subscriberid   uint `gorm:"index"`
	Besubscriberid uint
	Inuse          string
	Createtime     time.Time
}

func (*Subscribeuser) TableName() string {
	return "breakfast_subscribeuser"
}

func GetSubscribeById(db *gorm.DB, id uint) (*Subscribeuser, e.Error) {
	a := &Subscribeuser{}

	if err := db.Where("id = ?", id).Find(a).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			logger.Errorf("No Admin record found by [id = %d] : %s", id, err)
			return nil, errors.DBNotFound
		} else {
			logger.Errorf("Error querying Admin by [id = %d] : %s", id, err)
			return nil, errors.ServerInternalErr
		}
	}
	return a, nil
}
