package amqp

import "log"

const (
	AlternateExchangeName = "discarded"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
