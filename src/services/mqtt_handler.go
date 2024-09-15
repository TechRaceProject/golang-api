package services

import (
	"api/src/models"
	"api/src/models/attributes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type MQTTHandler struct{}

func (h MQTTHandler) HandleMQTTRaceData(id string, column string, payload string) {
	connection := GetConnection()

	var race models.Race

	if connection.Where("id = ?", id).First(&race).RowsAffected == 0 {
		return
	}

	var columnToUpdate string = column
	var valueToUpdate interface{}
	payload = strings.TrimSpace(payload)

	if column == "message" {
		vehicleId := race.VehicleID
		h.HandleCreateVehiculeHistory(vehicleId, payload)
		return
	}

	switch column {
	case "distance_covered", "out_of_parcours", "collision_duration":
		// we are always expecting a float from the ESP32 because of the way we are sending the data
		payloadToFloat, err := strconv.ParseFloat(payload, 64)

		if err != nil {
			fmt.Printf("Error while converting payload to float for %s: %v\n", column, err)
			return
		}

		valueToUpdate = int(payloadToFloat)

	case "average_speed":
		// we are always expecting a float from the ESP32 because of the way we are sending the data
		payloadToFloat, err := strconv.ParseFloat(payload, 64)

		if err != nil {
			fmt.Printf("Error while converting payload to float for %s: %v\n", column, err)
			return
		}

		valueToUpdate = payloadToFloat

	case "status":
		valueToUpdate = payload

	default:
		return
	}

	if race.Status == "completed" {
		fmt.Printf("['mqtt_handler] an update was requested for a completed race: %d - ignoring\n", race.ID)
		return
	}

	if race.Status == "not_started" {
		connection.Model(&race).Update("status", "in_progress")

		if race.StartTime == nil {
			connection.Model(&race).Update("start_time", &attributes.CustomTime{
				Time: time.Now(),
			})
		}
	}

	if columnToUpdate == "status" {
		connection.Model(&race).Update("status", valueToUpdate)

		if valueToUpdate == "completed" {

			race.EndTime = &attributes.CustomTime{
				Time: time.Now(),
			}

			connection.Model(&race).Update("end_time", race.EndTime)
		}

		return
	}

	connection.Model(&race).Update(columnToUpdate, valueToUpdate)
}

func (h MQTTHandler) HandleCreateVehiculeHistory(vehicule_id uint, payload string) {
	connection := GetConnection()

	var vehicle models.Vehicle

	if connection.Where("id = ?", vehicule_id).First(&vehicle).RowsAffected == 0 {
		fmt.Printf("Vehicle with id %d not found\n", vehicule_id)

		return
	}

	VehicleHistory := models.VehicleHistory{
		VehicleID: vehicle.ID,
		Message:   payload,
	}

	err := connection.Create(&VehicleHistory)

	if err.Error != nil {
		fmt.Printf("Error while creating vehicle history: %v\n", err.Error)
	}
}
