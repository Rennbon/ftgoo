package tool

import (
	"log"
)

func CallRecover(fun func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	fun()
}
