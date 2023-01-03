package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"
	"worlder-test/app/infrastructure/mqtt"

	"worlder-test/app/core/models"
	"worlder-test/app/core/usecase"
	"worlder-test/app/infrastructure/database"
	"worlder-test/app/interface/api"
	"worlder-test/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"
)

func assignRouting(e *echo.Echo) {
	//assign middlewares
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      nil,
		ErrorMessage: "Timeout Service",
		Timeout:      1 * time.Minute,
	}))

	//api handlers
	apiEngine := e.Group(config.Config.RootURL + "/api")

	// db connection
	dbConnection := config.GetInstanceDb()
	dbRepo := database.NewDbConnection(dbConnection)

	// mqtt connection
	mqttConnectionSub := config.GetInstanceMqtt(os.Getenv("MQTT_CLIENT_SUB"))
	mqttConnectionPub := config.GetInstanceMqtt(os.Getenv("MQTT_CLIENT_PUB"))
	mqttRepoSub := mqtt.NewMqttHandler(mqttConnectionSub)
	mqttRepoPub := mqtt.NewMqttHandler(mqttConnectionPub)

	// sensor api
	sensorUc := usecase.NewSensorUsecase(dbRepo)
	sensorApi := api.NewSensorApi(sensorUc)
	sensors := apiEngine.Group("/sensors")
	sensors.GET("/list", sensorApi.ListSensor)
	sensors.GET("/:id", sensorApi.GetSensorDetail)
	sensors.POST("/create", sensorApi.InsertDataSensor)
	sensors.PUT("/:id", sensorApi.UpdateDataSensor)
	sensors.DELETE("/:id", sensorApi.DeleteDataSensor)

	// mqtt sensor
	mqttUcSub := usecase.NewSensorMqttUc(dbRepo, mqttRepoSub)
	mqttUcPub := usecase.NewSensorMqttUc(dbRepo, mqttRepoPub)
	go mqttUcSub.InsertMessageViaMqtt()
	go mqttUcPub.PublishMessageViaMqtt()
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})

	var migrate bool
	flag.BoolVar(&migrate, "migrate", true, "If migrate true")
	flag.Parse()

	if migrate {
		models.Migrate()
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      AllowOriginSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders: []string{"*"},
	}))
	assignRouting(e)

	e.Server.Addr = config.Config.Port
	graceful.ListenAndServe(e.Server, 10*time.Second)
	logrus.Infoln("Start server on port : ", e.Server.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func AllowOriginSkipper(c echo.Context) bool {
	return false
}
