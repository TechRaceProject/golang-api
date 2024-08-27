package services

import (
	"fmt"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func messagePubHandler(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())

	// fmt.Println("############### MESSAGE RECEIVED ###############")
	// fmt.Println(msg.Topic())
	// fmt.Println(payload)
	// fmt.Println("############# END #################")

	messageParts := strings.Split(topic, "/")

	if messageParts[0] != "esp32" {
		return
	}

	if len(messageParts) < 4 {
		fmt.Println("Invalid message received. Must have at least 4 parts.")
		return
	}

	mqttHandler := MQTTHandler{}

	model := messageParts[1]
	id := messageParts[2]
	column := messageParts[3]

	switch model {
	case "races":
		mqttHandler.HandleMQTTRaceData(id, column, payload)
	default:
		return
	}
}

func connectHandler(client mqtt.Client) {
	fmt.Println("Connected")
}

func connectLostHandler(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func InitMQTTClient(brokerURL string) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURL)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := mqtt.NewClient(opts)
	token := client.Connect()

	if token.Wait() && token.Error() != nil {
		fmt.Println("MQTT client is not connected. Error:", token.Error())
	} else {
		fmt.Println("MQTT client is connected.")
	}

	return client
}

func Subscribe(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, messagePubHandler)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)
}
