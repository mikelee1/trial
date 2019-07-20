package main

import (
	"fmt"

	user2 "breakfast/services/user_center/models/user"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"myproj/try/testpostgres/util"
	"sync"
)

var wg2 = &sync.WaitGroup{}
var crnum = 1

func main() {
	db := util.CreateConn()
	defer db.Close()
	cli(db)
	//var i = 0
	//
	//wg2.Add(crnum)
	//t := time.Now()
	//for {
	//	if i>crnum-1{
	//		break
	//	}
	//	i++
	//	go func(db *gorm.DB,ii int,wg *sync.WaitGroup) {
	//		cli1(db,wg)
	//	}(db,i,wg2)
	//}
	//wg2.Wait()
	//fmt.Println(time.Now().Sub(t).Seconds())

	//cli2(db)

}


func cli(db *gorm.DB) {
	outUser := &user2.User{}
	//order.Order{}
	//order.OrderDanpin{}
	if err := db.Raw("select * FROM breakfast_order WHERE orderor = ?",18).Scan(outUser).Error;err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(outUser)

}


func cli1(db *gorm.DB,wg *sync.WaitGroup) {

	//defer db.Close()
	defer wg.Done()
	//if err := tx.Exec("lock table breakfast_user in EXCLUSIVE mode;").Error; err != nil {
	//	tx.Rollback()
	//}
	//tmpUser := &user2.User{Openid: openid}
	outUser := &user2.User{}
	renf := db.First(&user2.User{}, "id = ?", 18).Find(outUser).RecordNotFound()

	if !renf {
		fmt.Println("already exist")
		return
	}
	//err := tx.Model(user2.User{}).FirstOrCreate(outUser, tmpUser).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func cli2(db *gorm.DB) {

	sqlstr := ""
	var i = 0
	for {
		if i > 100 {
			break
		}
		openid := "test-" + uuid.New().String()
		renf := db.First(&user2.User{}, "openid = ?", openid).RecordNotFound()

		if !renf {
			fmt.Println("already exist")
		}

		if i == 0 {
			sqlstr += fmt.Sprintf(" (%s) ", openid)
		} else {
			sqlstr += fmt.Sprintf(" ,(%s) ", openid)
		}
		i++

	}
	err := db.Model(user2.User{}).Exec("insert into breakfast_user (openid) values " + sqlstr).Error
	if err != nil {
		fmt.Println(err)
	}
}
