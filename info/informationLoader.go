package info

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hendrik-scholz/iot-device-information-go/logging"
)

type authorization struct {
	Name      string `json:"name"`
	Role      string `json:"role"`
	DeedOwner string `json:"deedOwner"`
}

type geoPosition struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type scheduledTask struct {
	DateTime    string `json:"dateTime"`
	Description string `json:"description"`
}

type identification struct {
	Company        string          `json:"company"`
	Device         string          `json:"device"`
	Version        string          `json:"version"`
	ScheduledTasks []scheduledTask `json:"scheduledTasks"`
}

type uuid struct {
	UUID string `json:"uuid"`
}

func GetAuthorization() authorization {
	authorizationFromFile, error := ioutil.ReadFile("authorization.json")

	if error != nil {
		logging.GetLogger().Error(error.Error())
	}

	var deviceAuthorization authorization
	json.Unmarshal(authorizationFromFile, &deviceAuthorization)
	return deviceAuthorization
}

func GetGeoPosition() geoPosition {
	geoPositionFromFile, error := ioutil.ReadFile("geoposition.json")

	if error != nil {
		logging.GetLogger().Error(error.Error())
	}

	var deviceGeoPosition geoPosition
	json.Unmarshal(geoPositionFromFile, &deviceGeoPosition)
	return deviceGeoPosition
}

func GetIdentification() identification {
	identificationFromFile, error := ioutil.ReadFile("identification.json")

	if error != nil {
		logging.GetLogger().Error(error.Error())
	}

	var deviceIdentification identification
	json.Unmarshal(identificationFromFile, &deviceIdentification)
	return deviceIdentification
}

func GetUUID() uuid {
	uuidFromFile, error := ioutil.ReadFile("uuid.json")

	if error != nil {
		logging.GetLogger().Error(error.Error())
	}

	var deviceUuid uuid
	json.Unmarshal(uuidFromFile,
		&deviceUuid)
	return deviceUuid
}
