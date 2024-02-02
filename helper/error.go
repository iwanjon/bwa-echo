package helper

import (
	"log"
)

func PanicIfError(err error, s string) {
	if err != nil {
		log.Println(err)
		log.Println(s)
		panic(err)
	}
}
