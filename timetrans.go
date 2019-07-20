package main

import (
	"time"
	"strconv"
	"math"
	"fmt"
)

func main()  {
	t := time.Now().Add(-10*time.Hour)
	//fmt.Println(ueTime(t))
	fmt.Println(TranSelfTime(t))
}

func ueTime(t time.Time) string {

	//完整时间戳
	now := time.Now()
	d := now.Sub(t)
	s := d.Seconds()
	m := d.Minutes()
	h := d.Hours()
	fmt.Println(s,m,h)
	switch {
	case s<10:
		return "刚刚"
	case s<60:
		return "1分钟内"
	case m<60:
		m := strconv.Itoa(int(math.Floor(d.Minutes())))
		return m+"分钟前"
	case h<24:
		m := strconv.Itoa(int(math.Floor(d.Hours())))
		return m+"小时前"
	case h>24 && h<48:
		return "一天前"
	case h>48 && h<72:
		return "两天前"
	default:
		return t.Format("2006-01-02 15:04:05")
	}
}


func TranSelfTime(t time.Time) (string,string,string) {

	//完整时间戳
	now := time.Now()
	d := now.Sub(t)
	s := d.Seconds()
	m := d.Minutes()
	h := d.Hours()
	fmt.Println(s,m,h)


	today := now.Day()
	var rday,rnoon,rtime string
	switch {

	case h<24 && today == t.Day():
		rday = "今天"
	case h<48 && (today == t.Day() + 1||(today == 1)):
		rday = "昨天"
	default:
		rday = t.Format("2006-01-02 15:04:05")
	}

	creath := t.Hour()
	switch {
	case creath<6:
		rnoon = "凌晨"
	case creath<12:
		rnoon = "上午"
	case creath<18:
		rnoon = "下午"
	case creath<24:
		rnoon = "晚上"
	}
	rtime = t.Format("15:04")

	return rday,rnoon,rtime
}