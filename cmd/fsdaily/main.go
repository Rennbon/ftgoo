package main

import (
	"ftgoo/logic/mongodb"
	"time"
)

func main() {
	a := mongodb.FolderStatService{}
	a.DailyFlushing(time.Now().Local())
}
