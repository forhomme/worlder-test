package usecase

import (
	"fmt"
	"worlder-test/app/core/models"
	"worlder-test/app/core/ports"
)

type sensorUsecase struct {
	repo ports.DatabaseOutputPort
}

func NewSensorUsecase(repo ports.DatabaseOutputPort) ports.SensorInputPort {
	return &sensorUsecase{repo: repo}
}

func (s *sensorUsecase) ListSensors(req *models.RequestSensorParam) (resp *models.ResponseSensorData, err error) {
	data, count, err := s.repo.ListSensors(req)
	if err != nil {
		return
	}
	paging := models.Paging{
		Page:    req.Page,
		PerPage: req.PerPage,
		Counter: count,
	}
	resp = &models.ResponseSensorData{
		Data:   data,
		Paging: paging,
	}
	return
}

func (s *sensorUsecase) GetSensorByID(id string) (resp *models.Sensors, err error) {
	resp, err = s.repo.GetSensorByID(id)
	return
}

func (s *sensorUsecase) InsertSensor(in *models.Sensors) (resp *models.Sensors, err error) {
	resp, err = s.repo.InsertSensor(in)
	return
}

func (s *sensorUsecase) UpdateSensor(id string, in *models.Sensors) (resp *models.Sensors, err error) {
	existing, err := s.repo.GetSensorByID(id)
	if err != nil {
		return
	}
	if existing == nil {
		err = fmt.Errorf("Data ID not found : %v", id)
		return nil, err
	}
	resp, err = s.repo.UpdateSensor(id, in)
	return
}

func (s *sensorUsecase) DeleteSensor(id string) (err error) {
	existing, err := s.repo.GetSensorByID(id)
	if err != nil {
		return
	}
	if existing == nil {
		err = fmt.Errorf("Data ID not found : %v", id)
		return
	}
	err = s.repo.DeleteSensor(id)
	return
}
