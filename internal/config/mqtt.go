package config

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"sync"
)

type MqttConfig struct {
	client mqtt.Client
}

var (
	onceMqtt     sync.Once
	instanceMqtt *MqttConfig
)

func GetInstanceMqtt(client string) mqtt.Client {
	onceMqtt.Do(func() {
		mqttInfo := Config.Mqtt
		var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
			fmt.Println("Connected")
		}
		logs := fmt.Sprintf("[INFO] Connected to MQTT Host = %s ", mqttInfo.Host)

		opts := mqtt.NewClientOptions()
		opts.AddBroker(fmt.Sprintf("tcp://%s:%d", mqttInfo.Host, mqttInfo.Port))
		opts.SetClientID(client)
		opts.SetUsername(mqttInfo.Username)
		opts.SetPassword(mqttInfo.Password)
		opts.OnConnect = connectHandler

		c := mqtt.NewClient(opts)
		if token := c.Connect(); token.Wait() && token.Error() != nil {
			logs = fmt.Sprintf("[ERROR] Failed to connect to MQTT with err %s. Config=%s", token.Error(), mqttInfo.Host)
			log.Fatalln(logs)
		}
		fmt.Println(logs)
		instanceMqtt = &MqttConfig{client: c}
	})
	return instanceMqtt.client
}
