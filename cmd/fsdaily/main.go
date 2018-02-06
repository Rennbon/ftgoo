package main

import (
	"time"

	"github.com/Rennbon/ftgoo/logic/mongodb"
)

func main() {
	a := mongodb.FolderStatService{}
	a.DailyFlushing(time.Now().Local())
}
