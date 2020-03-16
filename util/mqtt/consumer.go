package mqtt

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Consumer struct {
	Client
}

// CreateConsumer create MQTT client for receiving message
func CreateConsumer(options Options) *Consumer {
	consumer := &Consumer{
		Client : connect("consumer", options.Topic)
	}

	m.Client.Subscribe(options.Topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})

	return consumer
}

func (m *Consumer) Consume(fn func([]byte)) {
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fn(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages from \"%s\" queue.")
	<-forever
}