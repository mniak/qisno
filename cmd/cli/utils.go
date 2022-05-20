package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

func handle(err error, messageAndArgs ...interface{}) {
	if err == nil {
		return
	}

	if len(messageAndArgs) == 0 {
		log.Fatalln(err)
		return
	}

	message := fmt.Sprint(messageAndArgs[0])
	if len(messageAndArgs) == 1 {
		log.Fatalln(errors.Wrap(err, message))
		return
	}

	tail := messageAndArgs[1:]
	message = fmt.Sprintf(message, tail...)
	log.Fatalln(errors.Wrap(err, message))
}
