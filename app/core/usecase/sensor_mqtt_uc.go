package usecase

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/labstack/gommon/log"
	"math/rand"
	"os"
	"time"
	"worlder-test/app/core/models"
	"worlder-test/app/core/ports"
)

type sensorMqttUc struct {
	dbRepo   ports.DatabaseOutputPort
	mqttRepo ports.MqttOutputPort
}

func NewSensorMqttUc(dbRepo ports.DatabaseOutputPort, mqttRepo ports.MqttOutputPort) ports.MqttInputPort {
	return &sensorMqttUc{
		dbRepo:   dbRepo,
		mqttRepo: mqttRepo,
	}
}

func (s *sensorMqttUc) InsertMessageViaMqtt() (err error) {
	for {
		s.mqttRepo.Subscribe(os.Getenv("MQTT_TOPIC"), func(client mqtt.Client, message mqtt.Message) {
			var data *models.Sensors
			err = json.Unmarshal(message.Payload(), &data)
			if err != nil {
				log.Errorf("Error unmarshal payload : ", err.Error())
				return
			}
			resp, err := s.dbRepo.InsertSensor(data)
			if err != nil {
				log.Errorf("Error Insert payload : ", err.Error())
				return
			}
			log.Info("Insert data success : ", resp.Base.ID)
			time.Sleep(1000 * time.Millisecond)
		})
	}
}

func (s *sensorMqttUc) PublishMessageViaMqtt() (err error) {
	for {
		randomValue := rand.Intn(100)
		randomId1 := rand.Intn(100)
		randomId2 := rand.Intn(100)
		data := models.Sensors{
			Value: fmt.Sprintf("%d", randomValue),
			ID1:   fmt.Sprintf("%d", randomId1),
			ID2:   fmt.Sprintf("%d", randomId2),
		}
		dataValue, errMarshal := json.Marshal(data)
		if errMarshal != nil {
			return
		}
		s.mqttRepo.Publish(os.Getenv("MQTT_TOPIC"), string(dataValue))
		time.Sleep(1 * time.Second)
	}
}
