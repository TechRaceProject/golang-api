package services

import (
    "log"
    "time"

    mqtt "github.com/eclipse/paho.mqtt.golang"
    "api/src/models"
    "gorm.io/gorm"
)

func MqttMessageHandler(db *gorm.DB) mqtt.MessageHandler {
    return func(client mqtt.Client, msg mqtt.Message) {
        topic := msg.Topic()
        payload := string(msg.Payload())

        sensorData := models.SensorData{
            Topic:     topic,
            Payload:   payload,
            Timestamp: time.Now(),
        }

        if err := db.Create(&sensorData).Error; err != nil {
            log.Printf("Erreur lors de l'insertion des données: %s", err)
        } else {
            log.Printf("Données insérées avec succès : %s", payload)
        }
    }
}

func InitMqttClient(broker string, db *gorm.DB) (mqtt.Client, error) {
    opts := mqtt.NewClientOptions()
    opts.AddBroker(broker)
    opts.SetClientID("go-mqtt-client")
    opts.SetDefaultPublishHandler(MqttMessageHandler(db))

    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        return nil, token.Error()
    }

    log.Println("Client MQTT connecté avec succès")

    topics := []string{"esp32/track", "esp32/sonar", "esp32/light"}
    for _, topic := range topics {
        if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
            log.Printf("Erreur lors de la souscription au topic %s: %s", topic, token.Error())
            return nil, token.Error()
        }
        log.Printf("Souscription au topic %s réussie", topic)
    }

    return client, nil
}
