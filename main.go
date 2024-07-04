package main

import (
	"api/src/routes"
	"encoding/json"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type SensorData struct {
    Light int    `json:"light"`
    Sonar int    `json:"sonar"`
    Track string `json:"track"`
}

func messagePubHandler(client mqtt.Client, msg mqtt.Message) {
    fmt.Println("Received message:")
		fmt.Println(string(msg.Payload()))
		fmt.Println(msg.Topic())

    var data SensorData
    err := json.Unmarshal(msg.Payload(), &data)
    if err != nil {
        log.Printf("Error unmarshalling JSON: %v", err)
        return
    }

    fmt.Printf("Light: %d, Sonar: %d, Track: %s\n", data.Light, data.Sonar, data.Track)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
    fmt.Printf("Connect lost: %v", err)
}

func main() {
	// fmt.Println("Starting server on port 8080...")
	fmt.Println("Starting here")


	// database, err := services.InitSqlConnection()

	// if err != nil {
	// 	log.Fatal("An error occurred with the database connection: ", err)
	// }

	// connection, err := database.DB()

	// if err != nil {
	// 	log.Fatal("An error occurred with the database connection: ", err)
	// }
	// defer connection.Close()

	// err = database.AutoMigrate(&models.User{})

	// if err != nil {
	// 	log.Fatal("Error performing database migrations: ", err)
	// }

	opts := mqtt.NewClientOptions()
	
	opts.AddBroker("tcp://mosquitto:1883")
	// opts.SetClientID("go_mqtt_client")
	// opts.SetUsername("YOUR_USERNAME") // si nécessaire
	// opts.SetPassword("YOUR_PASSWORD") // si nécessaire
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	// opts.AddBroker("")

	opts.SetDefaultPublishHandler(messagePubHandler)

	client := mqtt.NewClient(opts)
	token := client.Connect()

	if token.Wait() && token.Error() != nil {
		fmt.Println("MQTT client is not connected. Error:", token.Error())
	} else {
		fmt.Println("MQTT client is connected.")
	}

	subscribe(client)
		
	router := routes.SetupRouter()

	router.Run(":8000")

	// Maintenez la connexion ouverte
	select {}
}

func subscribe(client mqtt.Client) {
    topic := "esp32/#"
    token := client.Subscribe(topic, 1, messagePubHandler)
    token.Wait()
    fmt.Printf("Subscribed to topic: %s\n", topic)
}
