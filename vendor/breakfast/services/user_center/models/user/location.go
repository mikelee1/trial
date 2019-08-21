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

type Location struct {
	gorm.Model
	Yuanqu      int //index
	Building    int //几号楼,非index
	Floor       int //几层楼，非index
	CompName    string
	Dingcaner   int    //订餐人的id
	Qucanername string //取餐人可能和订餐人不一样
	Sex         bool
	Phone string
}

func (*Location) TableName() string {
	return "breakfast_user_location"
}

func GetLocationById(db *gorm.DB, id uint) (*Location, e.Error) {
	a := &Location{}

	if err := db.Model(&Location{}).Where("id = ?", id).Find(a).Error; err != nil {
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
