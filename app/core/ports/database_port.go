package ports

import (
	"worlder-test/app/core/models"
)

type DatabaseOutputPort interface {
	ListSensors(req *models.RequestSensorParam) (resp []models.Sensors, count int64, err error)
	GetSensorByID(id string) (resp *models.Sensors, err error)
	InsertSensor(in *models.Sensors) (resp *models.Sensors, err error)
	UpdateSensor(id string, in *models.Sensors) (resp *models.Sensors, err error)
	DeleteSensor(id string) (err error)
}
