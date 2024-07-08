package services

import (
	"api/src/models"
	"fmt"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func messagePubHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("Received message:")
	payload := string(msg.Payload())
	fmt.Println(payload)
	fmt.Println(msg.Topic())

	value, err := strconv.ParseFloat(payload, 64)
	if err != nil {
		fmt.Println("Error converting payload to float:", err)
		return
	}

	var sensorData models.SensorData

	switch msg.Topic() {
	case "esp32/track":
		sensorData.Track = value
	case "esp32/sonar":
		sensorData.Sonar = value
	case "esp32/light":
		sensorData.Light = value
	default:
		fmt.Println("Invalid topic")
		return
	}

	result := GetConnection().Create(&sensorData)
	if result.Error != nil {
		fmt.Println("Error inserting data into database:", result.Error)
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
