package main

import (
	"ftgoo/logic/mongodb"
)

func main() {
	a := mongodb.FolderStatService{}
	a.DailyFlushing()
}
