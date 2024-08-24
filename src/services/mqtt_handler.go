package services

import (
	"api/src/models"
	"fmt"
	"strconv"
)

type MQTTHandler struct{}

func (h MQTTHandler) HandleMQTTRaceData(id string, column string, payload string) {
	connection := GetConnection()

	var race models.Race

	if connection.Where("id = ?", id).First(&race).RowsAffected == 0 {
		return
	}

	var columnToUpdate string
	var valueToUpdate interface{}
	var err error

	convertToInt := func(payload string) (int, error) {
		return strconv.Atoi(payload)
	}

	convertToUint8 := func(payload string) (uint8, error) {
		uintValue, err := strconv.ParseUint(payload, 10, 8)
		return uint8(uintValue), err
	}

	switch column {
	case "distance_covered", "average_speed":
		columnToUpdate = column
		valueToUpdate, err = convertToInt(payload)
		if err != nil {
			fmt.Printf("Error while converting payload to int for %s: %v\n", column, err)
			return
		}

	case "number_of_collisions", "out_of_parcours":
		columnToUpdate = column
		valueToUpdate, err = convertToUint8(payload)
		if err != nil {
			fmt.Printf("Error while converting payload to uint8 for %s: %v\n", column, err)
			return
		}

	default:
		return
	}

	connection.Model(&race).Update(columnToUpdate, valueToUpdate)
}
