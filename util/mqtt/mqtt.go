package mqtt

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func connect(clientID string) mqtt.Client {
	opts := createClientOptions()
	client := mqtt.NewClient(opts)
	token := client.Connect()

	for !token.WaitTimeout(3 * time.Second) {}

	if err := token.Error(); err != nil {
		log.Fatal(err)
	}

	return client
}

func createClientOptions(clientID string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", os.Getenv("MQTT_HOST")))
	opts.SetUsername(os.Getenv("MQTT_USERNAME"))
	password, _ := os.Getenv("MQTT_PASSWORD")
	opts.SetPassword(password)
	opts.SetClientID(clientID)

	return opts
}