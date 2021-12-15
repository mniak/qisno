package main

import (
	"log"

	"github.com/mniak/pismo/domain"
)

var clock domain.ClockManager

func main() {
	a, err := initApplication()
	if err != nil {
		log.Fatalln(err)
	}
	err = a.runInterface()
	if err != nil {
		log.Fatalln(err)
	}
}
