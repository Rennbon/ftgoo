package main

import (
	"log"
	"time"

	cnf "github.com/Rennbon/ftgoo/config"

	"github.com/Rennbon/ftgoo/logic/mongodb"

	"github.com/robfig/cron"
)

var spec = "@daily"

func init() {
	conf, err := cnf.LoadConfig()
	if err != nil {
		panic(err)
	}
	if conf.Cronspec == "" {
		panic("Cronspec 读取失败")
	}
	spec = conf.Cronspec
}
func main() {
	log.Println(spec)
	c := cron.New()
	a := mongodb.FolderStatService{}
	//spec := "0  * * * *"
	c.AddFunc(spec, func() {
		log.Println("daily service has started")
		a.DailyFlushing(time.Now().Local())
		log.Println("daily service is completed")
	})
	c.Start()
	select {} //阻塞主线程不退出
}
