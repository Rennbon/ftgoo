package tool

import (
	"log"
)

func CallRecover() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from ", err)
		}
	}()
}
