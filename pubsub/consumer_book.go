package pubsub

import (
	"encoding/json"

	"github.com/rakhmadbudiono/simple-api/util/mqtt"
)

var books []Book

type ConsumerBook struct {
	consumer *mqtt.Consumer
}

func StartBookConsumer() {
	options := mqtt.Options{
		Connection: defaultMQTTConnection,
		Topic: "book",
	}

	cg := ConsumerBook{
		consumer: mqtt.CreateConsumer(options),
	}

	go cg.run()
}

func (m *ConsumerBook) run() {
	m.consumer.Consume(func(b []byte) {
		body := struct {
			BookID	uint32  `json:"id_book"`
			Title 	string 	`json:"title"`
		}{}

		if err := json.Unmarshal(b, &body); err != nil {
			failOnError(err, "fail to decode")
			return
		}

		book, err := book.New(body)

		if err != nil {
			failOnError(err, "fail to create book")
			return
		}

		books[len(books)] := book
	})
}