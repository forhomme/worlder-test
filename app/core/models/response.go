package models

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ResponseSensorData struct {
	Data   []Sensors `json:"data"`
	Paging Paging    `json:"paging"`
}

type ResponseMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(ctx echo.Context, response ResponseMessage) error {
	if response.Code == 0 {
		response.Code = http.StatusOK
	}

	if len(response.Message) == 0 {
		response.Message = "success"
	}

	ctx.Response().Header().Set(echo.HeaderServer, "Mysf Product Catalog/1.0 (arch64)")
	return ctx.JSON(response.Code, response)
}
