package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"worlder-test/app/core/models"
	"worlder-test/app/core/ports"
)

type SensorApi struct {
	repo ports.SensorInputPort
}

func NewSensorApi(repo ports.SensorInputPort) *SensorApi {
	return &SensorApi{repo: repo}
}

// ListSensor godoc
// @Summary List Data Sensor
// @Description List Data Sensor
// @Tags Sensors
// @Accept json
// @Produce json
// @Param id1   				query string false "ID 1"
// @Param id2 					query string false "ID 2"
// @Param duration   			query string false "Duration in Second"
// @Param time_from 			query string false "Time From data"
// @Param time_to          		query int    false "Time To data"
// @Param page                   query int    false "Page" default(1)
// @Param per_page               query int    false "Per Page" default(10)
// @Param order_by               query string false "Order By"
// @Param sort_type              query string false "Sort Type" default(asc)
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /engine/api/sensors/list [get]
func (s *SensorApi) ListSensor(ctx echo.Context) error {
	request := new(models.RequestSensorParam)
	err := ctx.Bind(request)
	if err != nil {
		return models.SendResponse(ctx, models.ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	resp, err := s.repo.ListSensors(request)
	if err != nil {
		return models.SendResponse(ctx, models.ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return models.SendResponse(ctx, models.ResponseMessage{
		Code:    http.StatusOK,
		Message: "success",
		Data:    resp,
	})
}

// GetSensorDetail godoc
// @Summary Get Detail Data Sensor
// @Description GEt Detail Data Sensor
// @Tags Sensors
// @Accept json
// @Produce json
// @Param id 	path string true "Data Sensor ID"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /engine/api/sensors/{id} [get]
func (s *SensorApi) GetSensorDetail(ctx echo.Context) error {
	id := ctx.Param("id")
	resp, err := s.repo.GetSensorByID(id)
	if err != nil {
		return models.SendResponse(ctx, models.ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return models.SendResponse(ctx, models.ResponseMessage{
		Code:    http.StatusOK,
		Message: "success",
		Data:    resp,
	})
}

// InsertDataSensor godoc
// @Summary Insert Data Sensor
// @Description Insert Data Sensor
// @Tags Sensors
// @Accept json
// @Produce json
// @Param request body models.Sensors true "Body Request"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /engine/api/sensors/create [post]
func (s *SensorApi) InsertDataSensor(ctx echo.Context) error {
	request := new(models.Sensors)
	err := ctx.Bind(request)
	if err != nil {
		return models.SendResponse(ctx, models.ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	resp, err := s.repo.InsertSensor(request)
	if err != nil {
		return models.SendResponse(ctx, models.ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return models.SendResponse(ctx, models.ResponseMessage{
		Code:    http.StatusOK,
		Message: "success",
		Data:    resp,
	})
}

// UpdateDataSensor godoc
// @Summary Update Data Sensor
// @Description Update Data Sensor
// @Tags Sensors
// @Accept json
// @Produce json
// @Param id 	path string true "Data Sensor ID"
// @Param 		request body models.Sensors true "Body Request"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /engine/api/sensors/{id} [put]
func (s *SensorApi) UpdateDataSensor(ctx echo.Context) error {
	request := new(models.Sensors)
	id := ctx.Param("id")
	err := ctx.Bind(request)
	if err != nil {
		return models.SendResponse(ctx, models.ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	resp, err := s.repo.UpdateSensor(id, request)
	if err != nil {
		return models.SendResponse(ctx, models.ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return models.SendResponse(ctx, models.ResponseMessage{
		Code:    http.StatusOK,
		Message: "success",
		Data:    resp,
	})
}

// DeleteDataSensor godoc
// @Summary Delete Data Sensor
// @Description Delete Data Sensor
// @Tags Sensors
// @Accept json
// @Produce json
// @Param id 	path string true "Data Sensor ID"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /engine/api/sensors/{id} [delete]
func (s *SensorApi) DeleteDataSensor(ctx echo.Context) error {
	id := ctx.Param("id")
	err := s.repo.DeleteSensor(id)
	if err != nil {
		return models.SendResponse(ctx, models.ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return models.SendResponse(ctx, models.ResponseMessage{
		Code:    http.StatusOK,
		Message: "success",
	})
}
