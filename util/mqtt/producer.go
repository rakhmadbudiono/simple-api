package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Producer struct {
	Client
	Topic string 
}

// CreateProducer create MQTT client for receiving message
func CreateProducer(options Options) *Producer {
	producer := &Producer{
		Client : connect("producer", options.Topic)
		Topic : options.Topic
	}

	return producer
}

func (m *Producer) Publish(message string) error {
	err := m.Client.Publish(m.Topic, 0, false, message)

	return err
}
