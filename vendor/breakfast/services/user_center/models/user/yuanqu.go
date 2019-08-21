package user

import (
	e "breakfast/common/errors"
	"github.com/jinzhu/gorm"

	"breakfast/common/errors"
	"github.com/op/go-logging"
)

func init() {
	logger = logging.MustGetLogger("user")
}

type Yuanqu struct {
	gorm.Model
	Name    string
	Address string
}

func (*Yuanqu) TableName() string {
	return "breakfast_yuanqu"
}

func GetYuanquById(db *gorm.DB, id uint) (*Yuanqu, e.Error) {
	a := &Yuanqu{}

	if err := db.Model(&Yuanqu{}).Where("id = ?", id).Find(a).Error; err != nil {
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
