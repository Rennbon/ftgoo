package main

import (
	"log"
	"time"

	"github.com/Rennbon/ftgoo/logic/mongodb"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	a := mongodb.FolderStatService{}
	spec := "@daily"
	//spec := "20 17 * * * *"
	c.AddFunc(spec, func() {
		log.Println("daily service has started")
		a.DailyFlushing(time.Now().Local())
		log.Println("daily service is completed")
	})
	c.Start()
	select {} //阻塞主线程不退出
}
