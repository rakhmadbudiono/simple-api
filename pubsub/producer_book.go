package pubsub

import (
	"encoding/json"

	"gitlab.com/eduka-omega/grader-service/util/mqtt"
)

type ProducerBook struct {
	producer *mqtt.Producer
}

var pb ProducerBook

func init() {
	options := mqtt.Options{
		Connection: defaultMQTTConnection,
		Topic: "book"
	}

	pb.producer := mqtt.CreateProducer(options)
}

func PublishBook(b Book) {
	pb.producer.publish(b)
}

func (m *ProducerScores) publish(b Book) {
	mBook, err := json.Marshal(b)
	if err != nil {
		failOnError(err, "failed to publish book")
		return
	}

	err = m.Publish(mBook)
	if err != nil {
		failOnError(err, "failed when publishing book")
	}
}