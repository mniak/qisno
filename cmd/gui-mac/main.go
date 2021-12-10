package main

import (
	"log"

	"github.com/mniak/qisno/pkg/qisno"
)

var clock qisno.ClockManager

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
