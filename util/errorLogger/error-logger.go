package errorLogger

import "log"

func FailOnError(err error, msg string) {
	log.Fatalf("%s: %s", msg, err)
}

func LogOnError(err error, msg string) {
	log.Printf("%s: %s", msg, err)
}

