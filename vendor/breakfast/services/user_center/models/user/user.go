package user

import (
	e "breakfast/common/errors"
	"github.com/jinzhu/gorm"

	"breakfast/common/errors"
	"github.com/op/go-logging"
)

var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("user")
}

type User struct {
	gorm.Model
	Openid   string  `gorm:"index"`
	Nickname string  `gorm:"default:'路人甲'"`
	Avatar   string  `gorm:"default:'null'"`
	Punish   string  `gorm:"default:'0'"`
	Sex      SexType `gorm:"default:'3'"`
	Inuse    string  `gorm:"default:'1'"`
	Phone    string
}

func (*User) TableName() string {
	return "breakfast_user"
}

func GetUserById(db *gorm.DB, id uint) (*User, e.Error) {
	a := &User{}

	if err := db.Model(&User{}).Where("id = ?", id).Find(a).Error; err != nil {
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

func GetUserByOpenid(db *gorm.DB, openid string) (*User, e.Error) {
	a := &User{}

	if err := db.Model(&User{}).Where("openid = ?", openid).Find(a).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			logger.Errorf("No Admin record found by [openid = %s] : %s", openid, err)
			return nil, errors.DBNotFound
		} else {
			logger.Errorf("Error querying Admin by [openid = %d] : %s", openid, err)
			return nil, errors.ServerInternalErr
		}
	}
	return a, nil
}
