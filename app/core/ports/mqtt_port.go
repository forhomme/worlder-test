package ports

import mqtt "github.com/eclipse/paho.mqtt.golang"

type MqttOutputPort interface {
	Publish(topic, message string)
	Subscribe(topic string, fn mqtt.MessageHandler)
}

type MqttInputPort interface {
	InsertMessageViaMqtt() (err error)
	PublishMessageViaMqtt() (err error)
}
