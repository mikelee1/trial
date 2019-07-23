package main

import (
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	i := 0
	c := cron.New()
	spec := "0 13 20 * * ?"
	c.AddFunc(spec, func() {
		i++
		//log.Println("cron running:", i)
		log.Println(time.Now())
	})
	c.Start()

	select {}
}
