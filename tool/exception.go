package tool

import (
	"log"
)

func CallRecover() {
	if err := recover(); err != nil {
		log.Println("recover from ", err)
	}
}
