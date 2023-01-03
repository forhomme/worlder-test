package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"worlder-test/app/core/ports"
)

type mqttHandler struct {
	client mqtt.Client
}

func NewMqttHandler(client mqtt.Client) ports.MqttOutputPort {
	return &mqttHandler{client: client}
}

func (m *mqttHandler) Publish(topic, message string) {
	token := m.client.Publish(topic, 0, false, message)
	token.Wait()
	return
}

func (m *mqttHandler) Subscribe(topic string, fn mqtt.MessageHandler) {
	token := m.client.Subscribe(topic, 1, fn)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)
	return
}
