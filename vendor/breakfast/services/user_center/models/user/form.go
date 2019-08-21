package user

import (
	e "breakfast/common/errors"
	"github.com/jinzhu/gorm"

	"breakfast/common/errors"
	"time"
)

type Form struct {
	gorm.Model
	Userid     int32 `gorm:"index"`
	Formid     string
	Formtype   string `gorm:"default:'ask'"`
	Createtime time.Time
	Inuse      string `gorm:"default:'1'"`
}

func (*Form) TableName() string {
	return "breakfast_form"
}

func GetFormById(db *gorm.DB, id uint) (*Form, e.Error) {
	a := &Form{}

	if err := db.Where("formid = ?", id).Find(a).Error; err != nil {
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
