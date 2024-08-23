package services

import (
	"fmt"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func messagePubHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("Received message:")
	payload := string(msg.Payload())
	fmt.Println(payload)
	fmt.Println(msg.Topic())

	_, err := strconv.ParseFloat(payload, 64)
	if err != nil {
		fmt.Println("Error converting payload to float:", err)
		return
	}

	switch msg.Topic() {
	case "esp32/track":
		fmt.Println("esp32/track")
	case "esp32/sonar":
		fmt.Println("esp32/sonar")
	case "esp32/light":
		fmt.Println("esp32/light")
	default:
		fmt.Println("Invalid topic")
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
