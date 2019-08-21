package user

import (
	"breakfast/common/errors"
	"github.com/jinzhu/gorm"
)

type OrderExp struct {
	gorm.Model
	Userid int32 `gorm:"index"`

	StartDateStr   string
	EndDateStr     string
	IsWorking      int32
	Position       string
	WorkCompany    string
	WorkDepartment string
	Description    string
	Inuse          string `gorm:"default:'1'"`
}

func (*OrderExp) TableName() string {
	return "breakfast_orderexp"
}

func GetOrderExpById(db *gorm.DB, id uint) (*OrderExp, errors.Error) {
	a := &OrderExp{}

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
