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

type Building struct {
	gorm.Model
	Name     int32 //1,2,3...
	Floornum int32 //一共多少层
	Yuanqu   uint  //隶属的园区
}

func (*Building) TableName() string {
	return "breakfast_building"
}

func GetBuildingById(db *gorm.DB, id uint) (*Building, e.Error) {
	a := &Building{}

	if err := db.Model(&Building{}).Where("id = ?", id).Find(a).Error; err != nil {
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
