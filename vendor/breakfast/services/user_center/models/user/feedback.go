package user

import (
	"breakfast/common/errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Feedback struct {
	gorm.Model
	Feedbackerid int32 `gorm:"index"`
	Content      string
	Connection   string
	Pic          string
	Createtime   time.Time
	Inuse        string `gorm:"default:'1'"`
}

func (*Feedback) TableName() string {
	return "breakfast_feedback"
}

func GetFeedbackById(db *gorm.DB, id uint) (*Feedback, errors.Error) {
	a := &Feedback{}

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
