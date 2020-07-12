package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"sync"
)

var (
	once *sync.Once
	Oconnect orm.Ormer
	dbtype = "mysql"
	dbname1 = "wasabi"
	dbuser = "yunphant"
	dbpasswd = "123456"
	dbip = "192.168.9.87"
	dbport = "38255"
	dbcharset = "utf8"
)

func init()  {
	once = &sync.Once{}

}

func CreateDBClient() {
	orm.RegisterModel(new(Auth))
	orm.RegisterModel(new(ChaincodeInfo))
	orm.RegisterModel(new(OrgChannel))
	orm.RegisterModel(new(Inform))
	orm.RegisterModel(new(Org))
	orm.RegisterModel(new(Node))
	connectstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", dbuser, dbpasswd, dbip, dbport, dbname1, dbcharset)
	// 数据库密码明文去除
	err := orm.RegisterDriver(dbtype, orm.DRMySQL)
	if err != nil {
		panic(err)
	}

	err = orm.RegisterDataBase("default", dbtype, connectstr)
	if err != nil {
		panic(err)
	}

	Oconnect = orm.NewOrm()
	orm.RunSyncdb("default", false, false)
}

func GetDBClient() orm.Ormer {
	once.Do(func() {
		CreateDBClient()
	})
	if Oconnect != nil{
		return Oconnect
	}
	return Oconnect
}
