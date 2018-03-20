package services

import (
	"log"
)

func IfErrorHandle(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
