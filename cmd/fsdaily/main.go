package main

import (
	"time"

	"github.com/Rennbon/ftgoo/logic/mongodb"
	"github.com/robfig/cron"
)

var refreshNow = true

func main() {
	c := cron.New()
	a := mongodb.FolderStatService{}
	c.AddFunc("@daily", func() {
		if refreshNow {
			a.DailyFlushing(time.Now().Local())
		} else {
			refreshNow = true
		}
	})
	c.Start()
	select {} //阻塞主线程不退出
}
