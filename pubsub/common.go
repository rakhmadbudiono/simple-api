package pubsub

import (
	"log"
	"os"

	"github.com/rakhmadbudiono/simple-api/util/mqtt"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

var defaultMQTTConnection mqtt.ConnectionOptions

func init() {
	defaultMQTTConnection.Username = os.Getenv("MQTT_USERNAME")
	defaultMQTTConnection.Password = os.Getenv("MQTT_PASSWORD")
	defaultMQTTConnection.Host = os.Getenv("MQTT_HOST")
	defaultMQTTConnection.Port = os.Getenv("MQTT_PORT")
	defaultClientIDConnection.ClientID = os.Getenv("MQTT_CLIENT_ID")
}
